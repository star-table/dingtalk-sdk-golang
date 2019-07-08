package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_FileUploadSingle(t *testing.T) {
	resp, _ := CreateClient().FileUploadSingle("C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_SendDingPanFileToSingleChat(t *testing.T) {
	//resp, err := CreateClient().SendDingPanFileToSingleChat()
	//t.Log(json.ToJson(resp), err)
}
