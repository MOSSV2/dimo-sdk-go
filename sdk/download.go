package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
	"sync"
	"time"

	bls "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
	"github.com/MOSSV2/dimo-sdk-go/lib/merkle"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/semaphore"
)

func DownloadPieceOrigin(baseUrl string, auth types.Auth, name string) ([]byte, error) {
	logger.Debug("download piece: ", name, " from: ", baseUrl)
	form := url.Values{}
	form.Set("name", name)
	form.Set("type", "piece")

	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, baseUrl, "/api/download", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	padbytes := bls.Pad(resByte)
	root, err := merkle.ReaderRoot(bytes.NewBuffer(padbytes), lighthash.New(), bls.PadSize)
	if err != nil {
		return nil, err
	}

	if strings.Compare(name, hex.EncodeToString(root)) != 0 {
		return nil, fmt.Errorf("invalid data for: %s", name)
	}

	return resByte, nil
}

func DownloadReplica(baseUrl string, auth types.Auth, piece, name string, addr common.Address, size int) ([]byte, error) {
	res, err := DownloadReplicaOrigin(baseUrl, auth, piece, name, addr, size)
	if err != nil {
		return DownloadReplicaFromStream(baseUrl, auth, piece, name, addr, size)
	}
	return res, nil
}

func DownloadReplicaFromStream(baseUrl string, auth types.Auth, piece, name string, addr common.Address, size int) ([]byte, error) {
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
		err = bls.SlothDecode(resByte, addr.Bytes())
		if err != nil {
			continue
		}

		root, err := merkle.ReaderRoot(bytes.NewBuffer(resByte), lighthash.New(), bls.PadSize)
		if err != nil {
			continue
		}

		if strings.Compare(piece, hex.EncodeToString(root)) != 0 {
			continue
		}

		resByte, err = bls.Unpad(resByte, size)
		if err != nil {
			continue
		}

		return resByte, nil
	}
	return nil, fmt.Errorf("fail")
}

func DownloadReplicaOrigin(baseUrl string, auth types.Auth, piece, name string, addr common.Address, size int) ([]byte, error) {
	logger.Debug("download replica: ", name, " from store")
	er, err := GetEdge(baseUrl, auth, addr)
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("name", name)

	logger.Debug("download replica: ", name, "from store: ", er.Name, " at:", er.ExposeURL)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, er.ExposeURL, "/api/download", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	err = bls.SlothDecode(resByte, addr.Bytes())
	if err != nil {
		return nil, err
	}

	root, err := merkle.ReaderRoot(bytes.NewBuffer(resByte), lighthash.New(), bls.PadSize)
	if err != nil {
		return nil, err
	}

	if strings.Compare(piece, hex.EncodeToString(root)) != 0 {
		return nil, fmt.Errorf("invalid data for: %s %s", piece, name)
	}

	resByte, err = bls.Unpad(resByte, size)
	if err != nil {
		return nil, err
	}

	return resByte, nil
}

func DownloadPiece(baseUrl string, auth types.Auth, name string) ([]byte, error) {
	logger.Debug("download: ", name, " from: ", baseUrl)
	pr, err := GetPieceReceipt(baseUrl, auth, name)
	if err != nil {
		return nil, err
	}
	if len(pr.Replicas) == 0 {
		url := baseUrl
		stream, err := GetEdge(baseUrl, auth, pr.StoredOn[0])
		if err == nil {
			url = stream.ExposeURL
		}

		return DownloadPieceOrigin(url, auth, name)
	}
	rmap := make(map[string]common.Address)
	for i, rep := range pr.Replicas {
		rmap[rep] = pr.StoredOn[i]
	}

	utils.ShuffleString(pr.Replicas)
	for _, rep := range pr.Replicas {
		val, err := DownloadReplicaFromStream(baseUrl, auth, name, rep, rmap[rep], int(pr.Size))
		if err != nil {
			logger.Debugf("fail download replica %s %s", rep, err)
			continue
		}

		return val, nil
	}

	return DownloadPieceOrigin(baseUrl, auth, name)
}

func DownloadPieceAndSave(baseUrl string, auth types.Auth, com string, ks types.IReplicaStore) error {
	if ks != nil {
		_, err := ks.Get(context.TODO(), com, nil, types.Options{})
		if err == nil {
			return nil
		}
	}

	resByte, err := DownloadPiece(baseUrl, auth, com)
	if err != nil {
		return err
	}

	if ks != nil {
		pr, err := GetPieceReceipt(baseUrl, auth, com)
		if err != nil {
			return err
		}
		ks.Put(context.TODO(), pr.PieceCore, resByte)
	}
	return nil
}

func CheckFile(baseUrl string, auth types.Auth, name string, ks types.IReplicaStore) error {
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	for _, com := range fr.Pieces {
		err = DownloadPieceAndSave(baseUrl, auth, com, ks)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckFileParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IReplicaStore) error {
	logger.Debug("download piece in parallel: ", parallel)
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(int64(parallel))
	for _, com := range fr.Pieces {
		err := sm.Acquire(context.TODO(), 1)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(com string, ks types.IReplicaStore) {
			defer sm.Release(1)
			defer wg.Done()

			DownloadPieceAndSave(baseUrl, auth, com, ks)
		}(com, ks)
	}
	wg.Wait()

	return nil
}

func Download(baseUrl string, auth types.Auth, name string, ks types.IReplicaStore, w io.Writer) error {
	fr, err := GetFileReceipt(baseUrl, auth, name)
	if err != nil {
		return err
	}

	h := sha256.New()
	restsize := fr.Size
	for _, com := range fr.Pieces {
		unpadsize := bls.MaxSize
		if restsize < bls.MaxSize {
			unpadsize = int(restsize)
		}

		if ks != nil {
			var b bytes.Buffer
			_, err := ks.Get(context.TODO(), com, &b, types.Options{
				UserDefined: map[string]string{
					"decode": "true",
				},
			})
			if err == nil {
				resByte := b.Bytes()

				resByte, err = bls.Unpad(resByte, unpadsize)
				if err != nil {
					return err
				}
				w.Write(resByte)
				h.Write(resByte)
				restsize -= bls.MaxSize
				continue
			}
		}

		resByte, err := DownloadPiece(baseUrl, auth, com)
		if err != nil {
			return err
		}

		if ks != nil {
			pr, err := GetPieceReceipt(baseUrl, auth, com)
			if err != nil {
				return err
			}
			ks.Put(context.TODO(), pr.PieceCore, resByte)
		}

		w.Write(resByte)
		h.Write(resByte)
		restsize -= bls.MaxSize
	}

	fh := h.Sum(nil)
	if fr.Name != hex.EncodeToString(fh) {
		return fmt.Errorf("invalid data for: %s", fr.Name)
	}

	return nil
}

func DownloadParallel(baseUrl string, auth types.Auth, name string, parallel int, ks types.IReplicaStore, w io.Writer) error {
	if ks != nil {
		err := CheckFileParallel(baseUrl, auth, name, parallel, ks)
		if err != nil {
			return err
		}
	}

	return Download(baseUrl, auth, name, ks, w)
}
