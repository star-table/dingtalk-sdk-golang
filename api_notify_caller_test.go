package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strings"
	"testing"
)

func GetTestInfo() (string, int, string) {
	corpTokenInfo, _ := GetCorpToken(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	accessToken := corpTokenInfo.AccessToken

	corpAuthInfo, _ := GetAuthInfo(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	agentId := corpAuthInfo.AuthInfo.Agent[0].AgentId

	memberList, _ := GetDepMember(accessToken, "1")
	userIdList := strings.Join(memberList.UserIds, ",")
	return accessToken, agentId, userIdList
}

func TestSendWorkTextNotice(t *testing.T) {
	accessToken, agentId, userIdList := GetTestInfo()
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

func TestSendWorkImageNotice(t *testing.T) {
	accessToken, agentId, userIdList := GetTestInfo()
	mediaResp, _ := UploadMedia(accessToken, "image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	//Image msg
	msg := WorkNoticeMsg{
		MsgType: "image",
		Image: &ImageNotice{
			MediaId: mediaResp.MediaId,
		},
	}
	resp, _ := SendWorkNotice(accessToken, agentId, &userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))
}

func TestSendWorkLinkNotice(t *testing.T) {
	accessToken, agentId, userIdList := GetTestInfo()
	mediaResp, _ := UploadMedia(accessToken, "image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	msg := WorkNoticeMsg{
		MsgType: "link",
		Link: &LinkNotice{
			MsgUrl: "http://blog.ikuvn.com",
			PicUrl: mediaResp.MediaId,
			Title:  "提醒喝水小助手",
			Text:   "该喝水了",
		},
	}
	resp, _ := SendWorkNotice(accessToken, agentId, &userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))
}
