package dingding_sdk_golang

import "testing"

func TestRobotSender(t *testing.T) {
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=8a306888cef34bd5c7c24a5677350d908aa81e9000cc1b7a1fbc73669e853a77"

	msg := WorkNoticeMsg{
		MsgType: "text",
		Text: &TextNotice{
			"我是本组织的提醒喝水小助手，记得喝水~",
		},
	}

	RobotSender(webHook, msg)

	client := CreateClient()
	mediaResp, _ := client.UploadMedia("image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	msg = WorkNoticeMsg{
		MsgType: "link",
		Link: &LinkNotice{
			MsgUrl: "http://blog.ikuvn.com",
			PicUrl: mediaResp.MediaId,
			Title:  "提醒喝水小助手",
			Text:   "该喝水了",
		},
	}

	RobotSender(webHook, msg)
}
