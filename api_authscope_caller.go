package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) GetAuthScopes() (GetAuthScopesResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}

	body, err := http.Get("https://oapi.dingtalk.com/auth/scopes", params)
	resp := GetAuthScopesResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
