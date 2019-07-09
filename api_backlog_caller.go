package dingding_sdk_golang

import (
	"fmt"
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) CreateOrUpdateBackLog(req SaveProcessRequest) (*CreateOrUpdateBackLogResp, error) {
	body, err := json.ToJson(req)
	fmt.Println(body)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"access_token":       client.AccessToken,
		"saveProcessRequest": encrypt.URLEncode(body),
	}

	respBody, err := http.Post("https://oapi.dingtalk.com/topapi/process/save", params, "")
	if err != nil {
		return nil, err
	}
	resp := &CreateOrUpdateBackLogResp{}
	json.FromJson(respBody, resp)
	return resp, err
}
