package sdk

import (
	"strings"
	"testing"

	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

func GetTestInfo() (string, *DingTalkClient) {
	client := CreateClient()
	memberList, _ := client.GetDepMemberIds("1")
	userIdList := strings.Join(memberList.UserIds, ",")
	return userIdList, client
}

func TestSendWorkTextNotice(t *testing.T) {
	userIdList, client := GetTestInfo()
	//Text msg
	msg := WorkNoticeMsg{
		MsgType: "text",
		Text: &TextNotice{
			"我是本组织的提醒喝水小助手，记得喝水8~",
		},
	}
	resp, _ := client.SendWorkNotice(&userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))
}

func TestSendWorkImageNotice(t *testing.T) {
	userIdList, client := GetTestInfo()
	mediaResp, _ := client.UploadMedia("image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	//Image msg
	msg := WorkNoticeMsg{
		MsgType: "image",
		Image: &ImageNotice{
			MediaId: mediaResp.MediaId,
		},
	}
	resp, _ := client.SendWorkNotice(&userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))
}

func TestSendWorkLinkNotice(t *testing.T) {
	userIdList, client := GetTestInfo()
	mediaResp, _ := client.UploadMedia("image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	msg := WorkNoticeMsg{
		MsgType: "link",
		Link: &LinkNotice{
			MsgUrl: "http://blog.ikuvn.com",
			PicUrl: mediaResp.MediaId,
			Title:  "提醒喝水小助手",
			Text:   "该喝水了",
		},
	}
	resp, _ := client.SendWorkNotice(&userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))
}

func TestGetWorkNoticeProgressAndResultAndRecall(t *testing.T) {
	userIdList, client := GetTestInfo()

	t.Log(userIdList)
	//Text msg
	msg := WorkNoticeMsg{
		MsgType: "text",
		Text: &TextNotice{
			"我是本组织的提醒喝水小助手，记得喝水4~",
		},
	}
	resp, _ := client.SendWorkNotice(&userIdList, nil, false, msg)
	t.Logf(json.ToJson(resp))

	progressResp, _ := client.GetWorkNoticeProgress(resp.TaskId)
	t.Logf(json.ToJson(progressResp))

	resultResp, _ := client.GetWorkNoticeSendResult(resp.TaskId)
	t.Logf(json.ToJson(resultResp))

	recallResp, _ := client.RecallWorkNotice(resp.TaskId)
	t.Logf(json.ToJson(recallResp))
}
