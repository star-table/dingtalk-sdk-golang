package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/encrypt"
	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
	"strconv"
)

//Desc: 工作通知消息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/is3zms
func (client *DingTalkClient) SendWorkNotice(userIdList *string, deptIdList *string, toAllUser bool, msg WorkNoticeMsg) (SendWorkNoticeResp, error) {
	msgJsonStr, _ := json.ToJson(msg)
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.FormatInt(client.AgentId, 10),
		"msg":          encrypt.URLEncode(msgJsonStr),
	}
	if userIdList != nil {
		params["userid_list"] = *userIdList
	}
	if deptIdList != nil {
		params["dept_id_list"] = *deptIdList
	}
	if toAllUser != false {
		params["to_all_user"] = "true"
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2", params, "")
	resp := SendWorkNoticeResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 查询工作通知消息的发送进度
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/is3zms#
func (client *DingTalkClient) GetWorkNoticeProgress(taskId int64) (GetWorkNoticeProgressResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.FormatInt(client.AgentId, 10),
		"task_id":      strconv.FormatInt(taskId, 10),
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/getsendprogress", params, "")
	resp := GetWorkNoticeProgressResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 查询工作通知消息的发送结果
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/is3zms#
func (client *DingTalkClient) GetWorkNoticeSendResult(taskId int64) (GetWorkNoticeResultResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.FormatInt(client.AgentId, 10),
		"task_id":      strconv.FormatInt(taskId, 10),
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/getsendresult", params, "")
	resp := GetWorkNoticeResultResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 工作通知消息撤回
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/is3zms#
func (client *DingTalkClient) RecallWorkNotice(taskId int64) (WorkNoticeRecallResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"agent_id":     strconv.FormatInt(client.AgentId, 10),
		"msg_task_id":  strconv.FormatInt(taskId, 10),
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/recall", params, "")
	resp := WorkNoticeRecallResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 发送普通消息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/wvdxel
func (client *DingTalkClient) SendNormalNotice(sender string, cid string, msg WorkNoticeMsg) (SendNormalNoticeResp, error) {
	msgJsonStr, _ := json.ToJson(msg)
	params := map[string]string{
		"access_token": client.AccessToken,
		"sender":       sender,
		"cid":          cid,
		"msg":          encrypt.URLEncode(msgJsonStr),
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/recall", params, "")
	resp := SendNormalNoticeResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
