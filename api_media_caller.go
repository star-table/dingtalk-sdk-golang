package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
)

func UploadMedia(accessToken string, mediaType string, path string) (UploadMediaResp, error) {

	params := map[string]string{
		"access_token": accessToken,
		"type":         mediaType,
	}
	body, err := http.PostFile("https://oapi.dingtalk.com/media/upload", params, path, "media")
	resp := UploadMediaResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
