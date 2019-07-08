package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_FileUploadSingle(t *testing.T) {
	resp, _ := CreateClient().FileUploadSingle("C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_SendDingPanFileToSingleChat(t *testing.T) {
	client := CreateClient()

	mediaId := "#iAEHAqRmaWxlA6h5dW5kaXNrMATOCw_n8wXNBvQGzWSTB85dIrQTCM0iXw"
	filename := "test.jpg"

	resp, _ := client.SendDingPanFileToSingleChat("15001956402427783", encrypt.URLEncode(mediaId), encrypt.URLEncode(filename))
	t.Log(json.ToJson(resp))
}
