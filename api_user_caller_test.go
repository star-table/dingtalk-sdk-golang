package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestGetDepMember(t *testing.T) {

	memberList, _ := CreateClient().GetDepMemberIds("1")
	t.Logf(json.ToJson(memberList))
}

func TestGetUserDetail(t *testing.T) {
	memberList, _ := CreateClient().GetDepMemberIds("1")
	t.Logf(json.ToJson(memberList))

	resp, _ := CreateClient().GetUserDetail(memberList.UserIds[2], nil)
	t.Logf(json.ToJson(resp))
}

func TestDingTalkClient_GetDepMemberList(t *testing.T) {
	resp, _ := CreateClient().GetDepMemberList("1", "", 1, 5, "")
	t.Logf(json.ToJson(resp))
}

func TestDingTalkClient_GetDepMemberDetailList(t *testing.T) {
	resp, _ := CreateClient().GetDepMemberDetailList("1", "", 1, 5, "")
	t.Logf(json.ToJson(resp))
}
