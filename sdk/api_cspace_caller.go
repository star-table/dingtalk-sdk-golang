package sdk

import (
	"fmt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"io"
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

//Desc: 分块上传文件：开启分块上传事务
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-12
func (client *DingTalkClient) BeginUploadTransaction(size int64, count int) (*BeginUploadTransaction, error) {
	params := map[string]string{
		"access_token":  client.AccessToken,
		"agent_id":      strconv.Itoa(client.AgentId),
		"file_size":     strconv.FormatInt(size, 10),
		"chunk_numbers": strconv.Itoa(count),
	}
	body, err := http.Post("https://oapi.dingtalk.com/file/upload/transaction", params, "")
	resp := BeginUploadTransaction{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 分块上传文件：上传文件块
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-12
func (client *DingTalkClient) BeginUploadChunk(uploadId string, sequence int, reader io.Reader) (*BaseResp, error) {
	params := map[string]string{
		"access_token":   client.AccessToken,
		"agent_id":       strconv.Itoa(client.AgentId),
		"upload_id":      uploadId,
		"chunk_sequence": strconv.Itoa(sequence),
	}
	body, err := http.PostFileWithReader("https://oapi.dingtalk.com/file/upload/chunk", params, reader)
	fmt.Println(body)
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 提交文件上传事务
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-12
func (client *DingTalkClient) CommitUploadTransaction(size int64, count int, uploadId string) (*CommitUploadTransaction, error) {
	params := map[string]string{
		"access_token":  client.AccessToken,
		"agent_id":      strconv.Itoa(client.AgentId),
		"upload_id":     uploadId,
		"file_size":     strconv.FormatInt(size, 10),
		"chunk_numbers": strconv.Itoa(count),
	}
	body, err := http.Get("https://oapi.dingtalk.com/file/upload/transaction", params)
	fmt.Println(body)
	resp := CommitUploadTransaction{}
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
func (client *DingTalkClient) AddFileToUserCSpace(code string, mediaId string, spaceId string, folderId string, name string, overwrite bool) (*AddFileToUserCSpaceResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.Itoa(client.AgentId),
		"code":         code,
		"media_id":     mediaId,
		"space_id":     spaceId,
		"folder_id":    folderId,
		"name":         name,
		"overwrite":    strconv.FormatBool(overwrite),
	}
	body, err := http.Post("https://oapi.dingtalk.com/cspace/add", params, "")
	resp := AddFileToUserCSpaceResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 获取企业下的自定义空间
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-7
func (client *DingTalkClient) GetCustomSpace(domain string) (*GetCustomSpaceResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
	}
	if client.AgentId != 0 {
		params["agent_id"] = strconv.Itoa(client.AgentId)
	}
	if domain != "" {
		params["domain"] = domain
	}
	body, err := http.Get("https://oapi.dingtalk.com/cspace/get_custom_space", params)
	resp := GetCustomSpaceResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//Desc: 授权用户访问企业自定义空间
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/zmmoa5#-8
func (client *DingTalkClient) GrantCustomSpace(domain string, grantType string, userId string, path string, fileIds string, duration int) (*BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"type":         grantType,
		"userid":       userId,
		"duration":     strconv.Itoa(duration),
	}
	if domain != "" {
		params["domain"] = domain
	}
	if path != "" {
		params["path"] = path
	}
	if fileIds != "" {
		params["fileIds"] = fileIds
	}
	if client.AgentId != 0 {
		params["agent_id"] = strconv.Itoa(client.AgentId)
	}
	if domain != "" {
		params["domain"] = domain
	}
	body, err := http.Get("https://oapi.dingtalk.com/cspace/grant_custom_space", params)
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}
