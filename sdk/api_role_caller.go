package sdk

import (
	"strconv"

	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

//Desc: 获取角色列表
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/qydf9c
func (client *DingTalkClient) GetRoleList(offset int, size int) (GetRoleListResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	if offset > -1 {
		params["offset"] = strconv.Itoa(offset)
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/role/list", params)
	resp := GetRoleListResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取角色下的员工列表
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/qydf9c
func (client *DingTalkClient) GetUsersInRole(roleId int64, offset int, size int) (GetUsersInRoleResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"role_id":      strconv.FormatInt(roleId, 10),
	}
	if offset > -1 {
		params["offset"] = strconv.Itoa(offset)
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/role/simplelist", params)
	resp := GetUsersInRoleResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取角色组
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/qydf9c
func (client *DingTalkClient) GetRoleGroup(groupId int64) (GetRoleGroupResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"group_id":     strconv.FormatInt(groupId, 10),
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/role/getrolegroup", params)
	resp := GetRoleGroupResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取角色详情
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/qydf9c
func (client *DingTalkClient) GetRoleDetail(roleId int64) (GetRoleDetailResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"roleId":       strconv.FormatInt(roleId, 10),
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/role/getrole", params)
	resp := GetRoleDetailResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
