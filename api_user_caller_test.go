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

func TestDingTalkClient_GetAdminList(t *testing.T) {
	resp, _ := CreateClient().GetAdminList()
	t.Logf(json.ToJson(resp))
}

func TestDingTalkClient_GetAdminScope(t *testing.T) {
	adminList, _ := CreateClient().GetAdminList()
	resp, _ := CreateClient().GetAdminScope(adminList.AdminList[0].UserId)
	t.Logf(json.ToJson(resp))
}

func TestDingTalkClient_CanAccessMicroApp(t *testing.T) {
	client := CreateClient()

	adminList, _ := client.GetAdminList()
	resp, _ := client.CanAccessMicroApp(adminList.AdminList[1].UserId, "22790")
	t.Logf(json.ToJson(resp))
}

func TestDingTalkClient_GetUserIdByUnionId(t *testing.T) {
	memberDetailList, _ := CreateClient().GetDepMemberDetailList("1", "", 0, 1, "")
	t.Log(json.ToJson(memberDetailList))
	resp, _ := CreateClient().GetUserIdByUnionId(memberDetailList.UserList[0].UnionId)
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_GetOrgUserCount(t *testing.T) {
	resp, _ := CreateClient().GetOrgUserCount(0)
	t.Log(json.ToJson(resp))

	resp, _ = CreateClient().GetOrgUserCount(1)
	t.Log(json.ToJson(resp))
}
