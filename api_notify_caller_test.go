package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strings"
	"testing"
)

func TestSendWorkNotice(t *testing.T) {

	corpTokenInfo, _ := GetCorpToken(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	accessToken := corpTokenInfo.AccessToken

	corpAuthInfo, _ := GetAuthInfo(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	agentId := corpAuthInfo.AuthInfo.Agent[0].AgentId

	memberList, _ := GetDepMember(accessToken, "1")
	userIdList := strings.Join(memberList.UserIds, ",")

	//Text msg
	msg := WorkNoticeMsg{
		MsgType: "text",
		Text: &TextNotice{
			"我是本组织的提醒喝水小助手，记得喝水~",
		},
	}
	resp, _ := SendWorkNotice(accessToken, agentId, &userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))

}
