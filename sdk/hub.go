package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func UploadHub(baseUrl string, mm types.MemeStruct) error {
	mmb, err := json.Marshal(mm)
	if err != nil {
		return err
	}

	resb, err := doRequest(context.TODO(), baseUrl, "/api/upload", "application/json", types.Auth{}, bytes.NewReader(mmb))
	if err != nil {
		return err
	}

	var res types.MemeMeta
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("upload done: ", res)

	return nil
}

func UploadHubData(baseUrl, owner, filename string, data []byte) error {
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

	resb, err := doRequest(context.TODO(), baseUrl, "/api/uploadData", mwriter.FormDataContentType(), types.Auth{}, ipr)
	if err != nil {
		return err
	}

	var res types.MemeMeta
	err = json.Unmarshal(resb, &res)
	if err != nil {
		return err
	}
	logger.Info("upload done: ", res)

	return nil
}

func DownloadHubData(baseUrl string, id, owner string) ([]byte, error) {
	form := url.Values{}
	form.Set("id", id)
	form.Set("owner", owner)

	logger.Debugf("download %s %s from hub %s", id, owner, baseUrl)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	resByte, err := doRequest(ctx, baseUrl, "/api/download", "", types.Auth{}, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}
