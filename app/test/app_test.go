package test

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/rand"
)

func doRequest(ctx context.Context, baseUrl, method, ctype string, r io.Reader) ([]byte, error) {
	haddr := baseUrl + method
	hreq, err := http.NewRequestWithContext(ctx, "POST", haddr, r)
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Content-Type", ctype)

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

	bar := progressbar.DefaultBytes(-1, baseUrl+" download:")

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := progressbar.NewReader(resp.Body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return res, nil
}

func UploadJson(baseUrl string, mm types.MemeStruct) error {
	mmb, err := json.Marshal(mm)
	if err != nil {
		return err
	}

	_, err = doRequest(context.TODO(), baseUrl, "/api/upload", "application/json", bytes.NewReader(mmb))
	if err != nil {
		return err
	}

	return nil
}

func TestUploadJson(t *testing.T) {
	url := "http://127.0.0.1:8086"

	for i := 0; i < 1; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		mm := types.MemeStruct{
			Owner:   "commonjson",
			ID:      hex.EncodeToString(nkey),
			Message: hex.EncodeToString(nval),
		}

		err := UploadJson(url, mm)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func UploadData(baseUrl, owner, filename string, data []byte) error {
	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	go func() {
		defer ipw.Close()
		defer mwriter.Close()

		err := mwriter.WriteField("owner", owner)
		if err != nil {
			return
		}

		part, err := mwriter.CreateFormFile("file", filename)
		if err != nil {
			return
		}

		part.Write(data)
	}()

	_, err := doRequest(context.TODO(), baseUrl, "/api/uploadData", mwriter.FormDataContentType(), ipr)
	if err != nil {
		return err
	}

	return nil
}

func TestUploadData(t *testing.T) {

	url := "http://127.0.0.1:8086"

	for i := 0; i < 1; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := UploadData(url, "common", hex.EncodeToString(nkey), nval)
		if err != nil {
			t.Fatal(err)
		}
	}
}
