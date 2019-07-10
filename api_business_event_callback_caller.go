package dingding_sdk_golang

import (
	"fmt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) RegisterCallBack(req RegisterCallBackReq) (*BaseResp, error) {
	reqJson, err := json.ToJson(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(reqJson)
	params := map[string]string{
		"access_token": client.AccessToken,
	}

	body, err := http.Post("https://oapi.dingtalk.com/call_back/register_call_back", params, reqJson)
	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	json.FromJson(body, resp)
	return resp, err
}
