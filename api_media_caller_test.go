package dingding_sdk_golang

import "testing"

func TestUploadMedia(t *testing.T) {
	resp, err := CreateClient().UploadMedia("image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	t.Log(resp, err)
}
