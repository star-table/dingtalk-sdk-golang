package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strconv"
)

//Desc: 工作通知消息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/is3zms
func SendWorkNotice(accessToken string, agentId int, userIdList *string, deptIdList *string, toAllUser bool, msg WorkNoticeMsg) (SendWorkNoticeResp, error) {
	msgJsonStr, _ := json.ToJson(msg)
	params := map[string]string{
		"access_token": accessToken,
		"agent_id":     strconv.Itoa(agentId),
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
