package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

//第三方企业应用免登
//https://open-doc.dingtalk.com/microapp/serverapi3/xcdh0r
func (client *DingTalkClient) GetUserInfo(code string) (GetUserInfoResp, error) {
	params := map[string]string{
		"access_token":client.AccessToken,
		"code":code,
	}

	body, err := http.Get("https://oapi.dingtalk.com/user/getuserinfo", params)
	resp := GetUserInfoResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}