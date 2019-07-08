package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"os"
	"strconv"
)

//Desc: 单步文件上传
//Doc: https://open-doc.dingtalk.com/microapp/serverapi2/wk3krc#g5zpqs
func (client *DingTalkClient) FileUploadSingle(path string) (*FileUploadSingleResp, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.Itoa(client.AgentId),
		"file_size":    strconv.FormatInt(fileInfo.Size(), 10),
	}
	body, err := http.PostFile("https://oapi.dingtalk.com/file/upload/single", params, path, "media")
	resp := FileUploadSingleResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 发送钉盘文件给指定用户
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-5
func (client *DingTalkClient) SendDingPanFileToSingleChat(userId string, mediaId string, fileName string) (*BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.Itoa(client.AgentId),
		"userid":       userId,
		"media_id":     mediaId,
		"file_name":    fileName,
	}
	body, err := http.Post("https://oapi.dingtalk.com/cspace/add_to_single_chat", params, "")
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 新增文件到用户钉盘
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-6
func (client *DingTalkClient) AddFileToUserCSpace(userId string, mediaId string, fileName string) (*BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.Itoa(client.AgentId),
		"userid":       userId,
		"media_id":     mediaId,
		"file_name":    fileName,
	}
	body, err := http.Post("https://oapi.dingtalk.com/cspace/add_to_single_chat", params, "")
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}
