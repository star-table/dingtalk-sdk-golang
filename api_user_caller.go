package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strconv"
)

//Desc: 获取用户详情
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
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

//Desc: 获取部门用户userid列表
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetDepMemberIds(deptId string) (GetDepMemberIdsResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"deptId":       deptId,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/getDeptMember", params)
	resp := GetDepMemberIdsResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

func (client *DingTalkClient) BaseGetDepMemberList(url string, deptId string, lang string, offset int64, size int64, order string) (GetDepMemberListResp, error) {
	params := map[string]string{
		"access_token":  client.AccessToken,
		"department_id": deptId,
	}
	if lang != "" {
		params["lang"] = lang
	}
	if offset > -1 && size > 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
		params["size"] = strconv.FormatInt(size, 10)
	}
	if order != "" {
		params["order"] = order
	}
	body, err := http.Get(url, params)
	resp := GetDepMemberListResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取部门用户
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetDepMemberList(deptId string, lang string, offset int64, size int64, order string) (GetDepMemberListResp, error) {
	return client.BaseGetDepMemberList("https://oapi.dingtalk.com/user/simplelist", deptId, lang, offset, size, order)
}

//Desc: 获取部门用户详情
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetDepMemberDetailList(deptId string, lang string, offset int64, size int64, order string) (GetDepMemberListResp, error) {
	return client.BaseGetDepMemberList("https://oapi.dingtalk.com/user/listbypage", deptId, lang, offset, size, order)
}

//Desc: 获取管理员列表
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetAdminList() (GetAdminListResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/get_admin", params)
	resp := GetAdminListResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取管理员通讯录权限范围
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetAdminScope(adminId string) (GetAdminScopeResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"userid":       adminId,
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/user/get_admin_scope", params)
	resp := GetAdminScopeResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 查询管理员是否具备管理某个应用的权限
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) CanAccessMicroApp(adminId string, appId string) (CanAccessMicroAppResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"userId":       adminId,
		"appId":        appId,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/can_access_microapp", params)
	resp := CanAccessMicroAppResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 根据unionid获取userid
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetUserIdByUnionId(unionId string) (GetUserIdByUnionIdResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"unionid":      unionId,
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/getUseridByUnionid", params)
	resp := GetUserIdByUnionIdResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取企业员工人数
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/momrx5
func (client *DingTalkClient) GetOrgUserCount(onlyActive int) (GetOrgUserCountResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"onlyActive":   strconv.Itoa(onlyActive),
	}
	body, err := http.Get("https://oapi.dingtalk.com/user/get_org_user_count", params)
	resp := GetOrgUserCountResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
