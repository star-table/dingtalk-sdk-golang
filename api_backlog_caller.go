package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

//Desc: 创建或更新待办模板
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/hhaed5
func (client *DingTalkClient) CreateOrUpdateBackLog(req SaveProcessRequest) (*CreateOrUpdateBackLogResp, error) {
	req.AgentId = client.AgentId
	body, err := json.ToJson(req)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"access_token":       client.AccessToken,
		"saveProcessRequest": encrypt.URLEncode(body),
	}

	respBody, err := http.Post("https://oapi.dingtalk.com/topapi/process/save", params, "")
	if err != nil {
		return nil, err
	}
	resp := &CreateOrUpdateBackLogResp{}
	json.FromJson(respBody, resp)
	return resp, err
}

//Desc: 删除待办模板
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/dvr6o6
func (client *DingTalkClient) DeleteBackLog(req DeleteBackLogReq) (*BaseResp, error) {
	req.AgentId = client.AgentId
	body, err := json.ToJson(req)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"access_token": client.AccessToken,
		"request":      encrypt.URLEncode(body),
	}

	respBody, err := http.Post("https://oapi.dingtalk.com/topapi/process/delete", params, "")
	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	json.FromJson(respBody, resp)
	return resp, err
}

//Desc: 创建待办实例
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/gk2b3e
func (client *DingTalkClient) CreateWorkRecord(req CreateWorkRecordRequest, title *string) (*CreateWorkRecordResp, error) {
	req.AgentId = client.AgentId
	reqJson, err := json.ToJson(req)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"access_token": client.AccessToken,
		"request":      encrypt.URLEncode(reqJson),
	}

	if title != nil {
		params["title"] = *title
	}

	respBody, err := http.Post("https://oapi.dingtalk.com/topapi/process/workrecord/create", params, "")
	if err != nil {
		return nil, err
	}
	resp := &CreateWorkRecordResp{}
	json.FromJson(respBody, resp)
	return resp, err
}

//Desc: 更新实例状态
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/vo0gqo
func (client *DingTalkClient) UpdateWorkRecord(req UpdateWorkRecordRequest) (*BaseResp, error) {
	req.AgentId = client.AgentId
	reqJson, err := json.ToJson(req)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"access_token": client.AccessToken,
		"request":      encrypt.URLEncode(reqJson),
	}

	respBody, err := http.Post("https://oapi.dingtalk.com/topapi/process/workrecord/update", params, "")

	if err != nil {
		return nil, err
	}
	resp := &BaseResp{}
	json.FromJson(respBody, resp)
	return resp, err
}
