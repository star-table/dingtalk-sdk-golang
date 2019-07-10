package sdk

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func (client *DingTalkClient) GetSubDept(id string) (GetSubdeptResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"id":           id,
	}

	body, err := http.Get("https://oapi.dingtalk.com/department/list_ids", params)
	resp := GetSubdeptResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetDeptList(lang *string, fetchChild *bool, id string) (GetDeptListResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"id":           id,
	}
	if lang != nil {
		params["lang"] = *lang
	}
	if *fetchChild != false {
		params["fetch_child"] = "true"
	}

	body, err := http.Get("https://oapi.dingtalk.com/department/list", params)
	resp := GetDeptListResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) GetDeptDetail(id string, lang *string) (GetDeptDetailResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"id":           id,
	}
	if lang != nil {
		params["lang"] = *lang
	}

	body, err := http.Get("https://oapi.dingtalk.com/department/get", params)
	resp := GetDeptDetailResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) ListParentDeptsByDept(id string) (ListParentDeptsByDeptResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"id":           id,
	}

	body, err := http.Get("https://oapi.dingtalk.com/department/list_parent_depts_by_dept", params)
	resp := ListParentDeptsByDeptResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) ListParentDepts(userId string) (ListParentDeptsResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"userId":       userId,
	}

	body, err := http.Get("https://oapi.dingtalk.com/department/list_parent_depts", params)
	resp := ListParentDeptsResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
