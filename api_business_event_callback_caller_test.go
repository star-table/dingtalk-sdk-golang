package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_RegisterCallBack(t *testing.T) {

	req := RegisterCallBackReq{
		CallBackTag: []string{"org_dept_remove "},
		Token:       "123456",
		AesKey:      "1234567890123456789012345678901234567890123",
		Url:         "http://cnico.vaiwan.com/test",
	}
	resp, err := CreateClient().RegisterCallBack(req)
	t.Log(json.ToJson(resp))
	t.Log(err)
}
