package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func Submit(baseUrl string, auth types.Auth, msm types.ServiceMeta) (types.ServiceMeta, error) {
	var re types.ServiceMeta
	mmb, err := json.Marshal(msm)
	if err != nil {
		return re, err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(res, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func ConfirmService(baseUrl string, auth types.Auth, sn, sroot string, pf []byte) error {
	form := url.Values{}
	form.Set("name", sn)
	form.Set("root", sroot)
	form.Set("proof", hex.EncodeToString(pf))

	_, err := doRequest(context.TODO(), baseUrl, "/api/confirmService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	return nil
}

func UpdateService(baseUrl string, auth types.Auth, sn string) error {
	form := url.Values{}
	form.Set("name", sn)

	_, err := doRequest(context.TODO(), baseUrl, "/api/updateService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	return nil
}

func GetService(baseUrl string, auth types.Auth, name string) (types.ServiceResult, error) {
	res := types.ServiceResult{}
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListService(baseUrl string, auth types.Auth, filter string) (types.ListServiceResult, error) {
	res := types.ListServiceResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listService", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("service list: ", res)
	return res, nil
}
