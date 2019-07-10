package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestGetUserInfoFromThird(t *testing.T) {
	resp, err := CreateClient().GetUserInfoFromThird("f7e89bd957463f41a99882cddb55a34a")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetUserInfoFromAdmin(t *testing.T) {
	resp, err := CreateClient().GetUserInfoFromAdmin("f95cf708e50630bf90c7221df8ed111b")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestDingTalkClient_GetUserInfoByCode(t *testing.T) {
	resp, err := CreateClient().GetUserInfoByCode("acd7dba7ae7b316a9c158dfdf7254ff6")
	t.Log(json.ToJson(resp))
	t.Log(err)
}
