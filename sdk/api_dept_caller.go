package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

//获取子部门ID列表
//https://open-doc.dingtalk.com/microapp/serverapi3/fuqv8x#hvdqc
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

//获取部门列表
//https://open-doc.dingtalk.com/microapp/serverapi3/fuqv8x#-1
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

//获取部门详情
//https://open-doc.dingtalk.com/microapp/serverapi3/fuqv8x#-2
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

//查询部门的所有上级父部门路径
//https://open-doc.dingtalk.com/microapp/serverapi3/fuqv8x#-3
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

//查询指定用户的所有上级父部门路径
//https://open-doc.dingtalk.com/microapp/serverapi3/fuqv8x#-4
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
