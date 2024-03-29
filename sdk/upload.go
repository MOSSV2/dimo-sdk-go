package sdk

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/archive"
	bls "github.com/MOSSV2/dimo-sdk-go/lib/bls12377"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

func Disorder(array []types.EdgeReceipt) {
	var temp types.EdgeReceipt
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(array) - 1; i >= 0; i-- {
		num := r.Intn(i + 1)
		temp = array[i]
		array[i] = array[num]
		array[num] = temp
	}
}

func UploadToStream(baseUrl string, auth types.Auth, filePath string) (types.FileReceipt, common.Address, error) {
	var res types.FileReceipt
	sinfo, err := Info(baseUrl)
	if err != nil {
		return res, common.Address{}, err
	}

	er, err := ListEdge(baseUrl, auth, types.StreamType)
	if err != nil {
		return res, common.Address{}, err
	}

	logger.Debug("streams before: ", er.Edges)
	Disorder(er.Edges)
	logger.Debug("streams after: ", er.Edges)

	if os.Getenv("STREAM_PRIORITY") != "" {
		//"0x3a2e98eaaa6ba9e3102cb622945f102c38221268"
		priaddr := common.HexToAddress(os.Getenv("STREAM_PRIORITY"))
		at := 0
		for j, em := range er.Edges {
			if em.Name == priaddr {
				at = j
				break
			}
		}

		if at != 0 {
			tmp := er.Edges[at]
			er.Edges[at] = er.Edges[0]
			er.Edges[0] = tmp
		}
	}

	for _, em := range er.Edges {
		if em.Type != types.StreamType {
			continue
		}
		fr, err := UploadData(em.ExposeURL, auth, filePath)
		if err != nil {
			logger.Debug("upload: ", filePath, " to: ", em.ExposeURL, " fail: ", err)
			if !strings.Contains(err.Error(), "already has file") {
				continue
			}
		}

		logger.Debug("upload meta: ", filePath, " to: ", baseUrl)
		res, err := UploadFileMeta(baseUrl, auth, em.Name, fr)
		return res, em.Name, err
	}

	fcws, err := UploadData(baseUrl, auth, filePath)
	if err != nil {
		return res, common.Address{}, err
	}

	res.FileCore = fcws.FileCore

	return res, sinfo.Name, nil
}

func UploadData(baseUrl string, auth types.Auth, filePath string) (types.FileCoreWithSize, error) {
	logger.Debug("upload: ", filePath, " to: ", baseUrl)
	var res types.FileCoreWithSize
	p, err := homedir.Expand(filePath)
	if err != nil {
		return res, err
	}

	bar := progressbar.DefaultBytes(-1, "upload:")

	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	go func() {
		defer ipw.Close()
		defer mwriter.Close()

		part, err := mwriter.CreateFormFile("file", p)
		if err != nil {
			return
		}
		pf, err := os.Open(p)
		if err != nil {
			return
		}
		defer pf.Close()
		pr := progressbar.NewReader(pf, bar)

		io.Copy(part, &pr)
	}()

	haddr := baseUrl + "/api/upload"
	hreq, err := http.NewRequest("POST", haddr, ipr)
	if err != nil {
		return res, err
	}

	aub, err := json.Marshal(auth)
	if err != nil {
		return res, err
	}

	hreq.Header.Add("Authorization", hex.EncodeToString(aub))
	hreq.Header.Add("Content-Type", mwriter.FormDataContentType())
	defaultHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			WriteBufferSize:       16 << 10, // 16KiB moving up from 4KiB default
			ReadBufferSize:        16 << 10, // 16KiB moving up from 4KiB default
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
		},
	}

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s, msg: %s", resp.Status, resByte)
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	bar.Finish()

	return res, nil
}

func walk(baseDir, curdir string) (map[string]string, uint64, error) {
	res := make(map[string]string)
	size := uint64(0)
	fulDir := path.Join(baseDir, curdir)
	rd, err := os.ReadDir(fulDir)
	if err != nil {
		return res, size, err
	}
	h := sha256.New()
	for _, fde := range rd {
		if fde.IsDir() {
			subCur := path.Join(curdir, fde.Name())
			subRes, subSize, err := walk(baseDir, subCur)
			for k, v := range subRes {
				res[k] = v
			}
			if err != nil {
				return res, size, err
			}
			size += subSize
			continue
		}

		fi, err := fde.Info()
		if err != nil {
			return res, size, err
		}
		if fi.Size() < bls.MaxSize && fi.Name() != archive.ShadowTar {
			continue
		}
		osf, err := os.OpenFile(path.Join(fulDir, fi.Name()), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return res, size, err
		}
		h.Reset()
		_, err = io.Copy(h, osf)
		if err != nil {
			osf.Close()
			return res, size, err
		}
		osf.Close()
		hs := h.Sum(nil)
		res[path.Join(curdir, fi.Name())] = hex.EncodeToString(hs)
		size += uint64(fi.Size())
	}
	return res, size, nil
}
