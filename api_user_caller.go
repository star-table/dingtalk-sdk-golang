package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func GetUserDetail(accessToken string, userId string, lang *string) (GetUserDetailResp, error) {
	params := map[string]string{
		"access_token": accessToken,
		"userid":       userId,
		"lang":         *lang,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/get", params)
	resp := GetUserDetailResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func GetDepMember(accessToken string, deptId string) (GetDepMemberResp, error) {
	params := map[string]string{
		"access_token": accessToken,
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
