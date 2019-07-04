package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"strconv"
)

func SendWorkNotice(accessToken string, agentId int, userIdList string, deptIdList string, toAllUser bool, msg string) (string, error) {
	params := map[string]string{
		"access_token": accessToken,
		"agent_id":     strconv.Itoa(agentId),
		"userid_list":  userIdList,
		"dept_id_list": deptIdList,
		"to_all_user":  strconv.FormatBool(toAllUser),
		"msg":          msg,
	}
	return http.Post("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2", params, "")
}
