package sdk

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

//第三方企业应用免登(获取用户userid)
//https://open-doc.dingtalk.com/microapp/serverapi3/xcdh0r
func (client *DingTalkClient) GetUserInfoFromThird(code string) (GetUserInfoFromThirdResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"code":         code,
	}

	body, err := http.Get("https://oapi.dingtalk.com/user/getuserinfo", params)
	resp := GetUserInfoFromThirdResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//应用管理后台免登(获取应用管理员的身份信息)
//https://open-doc.dingtalk.com/microapp/serverapi3/ydc8us
func (client *DingTalkClient) GetUserInfoFromAdmin(code string) (GetUserInfoFromAdminResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"code":         code,
	}

	body, err := http.Get("https://oapi.dingtalk.com/sso/getuserinfo", params)
	resp := GetUserInfoFromAdminResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetUserInfoByCode(code string) (GetUserInfoByCodeResp, error) {
	params := map[string]string{
		"access_token":  client.AccessToken,
		"tmp_auth_code": code,
	}

	body, err := http.Post("https://oapi.dingtalk.com/sns/getuserinfo_bycode", params, "")
	resp := GetUserInfoByCodeResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
