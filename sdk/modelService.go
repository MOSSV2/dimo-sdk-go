package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func SubmitModelService(baseUrl string, auth types.Auth, msm types.ModelServiceMeta, pf []byte) (types.ModelServiceMeta, error) {
	var re types.ModelServiceMeta
	mmb, err := json.Marshal(msm)
	if err != nil {
		return re, err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	form.Set("proof", hex.EncodeToString(pf))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitModelService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(res, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func ConfirmModelService(baseUrl string, auth types.Auth, sn, sroot string, pf []byte) error {
	form := url.Values{}
	form.Set("name", sn)
	form.Set("root", sroot)
	form.Set("proof", hex.EncodeToString(pf))

	_, err := doRequest(context.TODO(), baseUrl, "/api/confirmModelService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	return nil
}

func GetModelService(baseUrl string, auth types.Auth, name string) (types.ModelServiceResult, error) {
	res := types.ModelServiceResult{}
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getModelService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListModelService(baseUrl string, auth types.Auth, filter string) (types.ListModelServiceResult, error) {
	res := types.ListModelServiceResult{}
	opt := types.Options{
		UserDefined: make(map[string]string),
	}
	opt.UserDefined["filter"] = filter

	optyByte, err := json.Marshal(opt)
	if err != nil {
		return res, err
	}

	form := url.Values{}
	form.Set("option", hex.EncodeToString(optyByte))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listModelService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("model service list: ", res)
	return res, nil
}
