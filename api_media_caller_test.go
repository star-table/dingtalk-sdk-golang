package dingding_sdk_golang

import "testing"

func TestUploadMedia(t *testing.T) {
	corpTokenInfo, _ := GetCorpToken(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	accessToken := corpTokenInfo.AccessToken
	resp, err := UploadMedia(accessToken, "image", "C:\\Users\\admin\\Desktop\\dingding-test.jpg")
	t.Log(resp, err)
}
