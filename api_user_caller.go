package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) GetUserDetail(userId string, lang *string) (GetUserDetailResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"userid":       userId,
	}
	if lang != nil {
		params["lang"] = *lang
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/get", params)
	resp := GetUserDetailResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetDepMember(deptId string) (GetDepMemberResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"deptId":       deptId,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/getDeptMember", params)
	resp := GetDepMemberResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
