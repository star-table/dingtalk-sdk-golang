package sdk

import (
	"github.com/polaris-team/dingtalk-sdk-golang/http"
	"github.com/polaris-team/dingtalk-sdk-golang/json"
)

//创建场景群
//https://developers.dingtalk.com/document/chatgroup/create-a-scene-group-v2?spm=ding_open_doc.document.0.0.1b9c70dfJlvfnF#topic-2019749
func (client *DingTalkClient) CreateSceneGroup(req CreateSceneGroupReq) (CreateSceneGroupResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	reqBodyJson, _ := json.ToJson(req)
	body, err := http.Post("https://oapi.dingtalk.com/topapi/im/chat/scenegroup/create", params, reqBodyJson)
	resp := CreateSceneGroupResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}

//更新场景群
//https://developers.dingtalk.com/document/chatgroup/scene-group-update
func (client *DingTalkClient) UpdateSceneGroup(req UpdateSceneGroupReq) (BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	reqBodyJson, _ := json.ToJson(req)
	body, err := http.Post("https://oapi.dingtalk.com/topapi/im/chat/scenegroup/update", params, reqBodyJson)
	resp := BaseResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}

//查询群信息
//https://developers.dingtalk.com/document/chatgroup/queries-the-basic-information-of-a-scenario-group
func (client *DingTalkClient) GetSceneGroup(req GetSceneGroupReq) (GetSceneGroupResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	reqBodyJson, _ := json.ToJson(req)
	body, err := http.Post("https://oapi.dingtalk.com/topapi/im/chat/scenegroup/get", params, reqBodyJson)
	resp := GetSceneGroupResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}

//新增群成员
//https://oapi.dingtalk.com/topapi/im/chat/scenegroup/member/add
func (client *DingTalkClient) AddSceneGroupMember(req AddSceneGroupMemberReq) (BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	reqBodyJson, _ := json.ToJson(req)
	body, err := http.Post("https://oapi.dingtalk.com/topapi/im/chat/scenegroup/member/add", params, reqBodyJson)
	resp := BaseResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}

//移除群成员
//https://oapi.dingtalk.com/topapi/im/chat/scenegroup/member/delete
func (client *DingTalkClient) DeleteSceneGroupMember(req DeleteSceneGroupMemberReq) (BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	reqBodyJson, _ := json.ToJson(req)
	body, err := http.Post("https://oapi.dingtalk.com/topapi/im/chat/scenegroup/member/delete", params, reqBodyJson)
	resp := BaseResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}
