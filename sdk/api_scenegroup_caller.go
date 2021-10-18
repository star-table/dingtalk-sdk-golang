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
	body, err := http.Post("https://oapi.dingtalk.com/topapi/v2/user/list", params, reqBodyJson)
	resp := CreateSceneGroupResp{}
	if err != nil {
		return resp, err
	}
	_ = json.FromJson(body, &resp)
	return resp, err
}
