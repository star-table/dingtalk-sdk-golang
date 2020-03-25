package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
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
