package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"

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
