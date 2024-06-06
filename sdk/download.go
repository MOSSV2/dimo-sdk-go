package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"
	"github.com/MOSSV2/dimo-sdk-go/lib/bls/erasure"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/semaphore"
)

func DownloadPiece(baseUrl string, auth types.Auth, name string, streamer common.Address) ([]byte, error) {
	logger.Debug("download piece: ", name, " from: ", baseUrl)
	pr, err := GetPieceReceipt(baseUrl, auth, name)
	if err != nil {
		return nil, err
	}

	res := make([][]byte, pr.Policy.N)
	suc := 0
	need := make([]int, 0, pr.Policy.K)
	survived := make([]int, 0, pr.Policy.K)
	for i, rep := range pr.Replicas {
		if rep == "" {
			logger.Warnf("piece %s %d is not stored", pr.Name, i)
			continue
		}
		suc++
	}
	if suc < int(pr.Policy.K) {
		er, err := GetEdge(baseUrl, auth, streamer)
		if err != nil {
			logger.Warnf("%s is not online: %s", streamer, err)
			return nil, err
		}
		baseUrl = er.ExposeURL
		pr, err = GetPieceReceipt(baseUrl, auth, name)
		if err != nil {
			logger.Warn("get piece: ", err)
			return nil, err
		}
	}

	suc = 0
	for i, rep := range pr.Replicas {
		val, err := DownloadReplica(baseUrl, auth, rep, pr.StoredOn[i])
		if err != nil {
			if i < int(pr.Policy.K) {
				need = append(need, i)
			}
			logger.Debugf("fail download replica %s %s", rep, err)
			continue
		}

		res[i] = val
		suc++
		survived = append(survived, i)
		if suc >= int(pr.Policy.K) {
			break
		}
	}
	if suc < int(pr.Policy.K) {
		return nil, fmt.Errorf("no enough replica")
	}
	// repair
	if len(need) > 0 {
		rs, err := erasure.NewRS(int(pr.Policy.N), int(pr.Policy.K))
		if err != nil {
			return nil, err
		}
		re, err := rs.NewReconst(survived)
		if err != nil {
			return nil, err
		}
		surdata := make([][]byte, pr.Policy.K)
		for i := 0; i < int(pr.Policy.K); i++ {
			surdata[i] = res[survived[i]]
		}
		ndata, err := re.Encode(surdata, need)
		if err != nil {
			return nil, err
		}
		for i, v := range need {
			res[v] = ndata[i]
		}
	}

	pbyte := make([]byte, 0, pr.Size)
	for i := 0; i < int(pr.Policy.K); i++ {
		res[i], err = bls.Unpad(res[i])
		if err != nil {
			return nil, err
		}
		pbyte = append(pbyte, res[i]...)
	}

	return pbyte[:pr.Size], nil
}

func DownloadReplica(baseUrl string, auth types.Auth, name string, addr common.Address) ([]byte, error) {
	res, err := DownloadReplicaOrigin(baseUrl, auth, name)
	if err != nil {
		return DownloadReplicaFromStream(baseUrl, auth, name, addr)
	}
	return res, nil
}

func DownloadReplicaOrigin(baseUrl string, auth types.Auth, name string) ([]byte, error) {
	form := url.Values{}
	form.Set("name", name)
	form.Set("type", "replica")

	logger.Debug("download replica: ", name, " at:", baseUrl)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, baseUrl, "/api/download", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}

func DownloadReplicaFromStream(baseUrl string, auth types.Auth, name string, addr common.Address) ([]byte, error) {
	logger.Debug("download replica: ", name, " from stream")
	el, err := ListEdge(baseUrl, auth, types.StreamType)
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("name", name)
	form.Set("storedOn", addr.String())

	for _, er := range el.Edges {
		logger.Debug("download replica: ", name, " from stream: ", er.Name, " at: ", er.ExposeURL)
		ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
		defer cancle()
		resByte, err := doRequest(ctx, er.ExposeURL, "/api/download", auth, strings.NewReader(form.Encode()))
		if err != nil {
			continue
		}

		return resByte, nil
	}
	return nil, fmt.Errorf("fail")
}

func DownloadPieceAndSave(baseUrl string, auth types.Auth, com string, ks types.IPieceStore, streamer common.Address) error {
	if ks != nil {
		_, err := ks.GetPiece(context.TODO(), com, nil, types.Options{})
		if err == nil {
			return nil
		}
	}

	resByte, err := DownloadPiece(baseUrl, auth, com, streamer)
	if err != nil {
		return err
	}

	if ks != nil {
		pr, err := GetPieceReceipt(baseUrl, auth, com)
		if err != nil {
			return err
		}
		ks.PutPiece(context.TODO(), pr.PieceCore, resByte, true)
	}
	return nil
}

func CheckFile(baseUrl string, auth types.Auth, name string, ks types.IPieceStore) error {
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	for i, com := range fr.Pieces {
		err = DownloadPieceAndSave(baseUrl, auth, com, ks, fr.Streams[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckFileParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore) error {
	logger.Debug("download piece in parallel: ", parallel)
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(int64(parallel))
	for i, com := range fr.Pieces {
		err := sm.Acquire(context.TODO(), 1)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(ni int, com string, ks types.IPieceStore) {
			defer sm.Release(1)
			defer wg.Done()

			DownloadPieceAndSave(baseUrl, auth, com, ks, fr.Streams[ni])
		}(i, com, ks)
	}
	wg.Wait()

	return nil
}

func Download(baseUrl string, auth types.Auth, name string, ks types.IPieceStore, w io.Writer) error {
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	for i, com := range fr.Pieces {
		if ks != nil {
			var b bytes.Buffer
			_, err := ks.GetPiece(context.TODO(), com, &b, types.Options{})
			if err == nil {
				w.Write(b.Bytes())
				continue
			}
		}

		resByte, err := DownloadPiece(baseUrl, auth, com, fr.Streams[i])
		if err != nil {
			return err
		}

		if ks != nil {
			pr, err := GetPieceReceipt(baseUrl, auth, com)
			if err != nil {
				return err
			}
			ks.PutPiece(context.TODO(), pr.PieceCore, resByte, true)
		}

		w.Write(resByte)
	}

	return nil
}

func DownloadParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IPieceStore, w io.Writer) error {
	if ks != nil {
		err := CheckFileParallel(baseUrl, auth, name, parallel, ks)
		if err != nil {
			return err
		}
	}

	return Download(baseUrl, auth, name, ks, w)
}

// todo: handle size > piece size
func DownloadWSize(baseUrl string, auth types.Auth, name string, ks types.IPieceStore, w io.Writer, start, size int64) error {
	logger.Debugf("download file %s %d %d ", name, start, size)
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	sstart := int64(0)
	pstart := int64(0)
	for i, com := range fr.Pieces {
		pr, err := GetPieceReceipt(baseUrl, auth, com)
		if err != nil {
			return err
		}

		tmp := sstart + pr.Size
		if start >= tmp {
			sstart = tmp
			continue
		}
		if sstart >= start+size {
			return nil
		}
		pstart = start - sstart
		sstart = tmp

		if ks != nil {
			var b bytes.Buffer
			_, err := ks.GetPiece(context.TODO(), com, &b, types.Options{})
			if err == nil {
				_, err = w.Write(b.Bytes())
				return err
			}
		}

		resByte, err := DownloadPiece(baseUrl, auth, com, fr.Streams[i])
		if err != nil {
			return err
		}

		if ks != nil {
			err = ks.PutPiece(context.TODO(), pr.PieceCore, resByte, true)
			if err != nil {
				return err
			}
			var b bytes.Buffer
			_, err = ks.GetPiece(context.TODO(), com, &b, types.Options{
				UserDefined: map[string]string{
					"start": strconv.FormatInt(pstart, 10),
					"size":  strconv.FormatInt(size, 10),
				},
			})
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			return err
		}
	}

	return nil
}
