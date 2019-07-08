package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) ListLabelGroups(size string, offset string) (ListLabelGroupsResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"size":         size,
		"offset":       offset,
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/extcontact/listlabelgroups", params, "")
	resp := ListLabelGroupsResp{}

	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetExtcontactList(size string, offset string) (GetExtcontactListResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"size":         size,
		"offset":       offset,
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/extcontact/list", params, "")
	resp := GetExtcontactListResp{}

	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetExtcontactDetail(userId string) (GetExtcontactDetailResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"user_id":      userId,
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/extcontact/get", params, "")
	resp := GetExtcontactDetailResp{}

	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
