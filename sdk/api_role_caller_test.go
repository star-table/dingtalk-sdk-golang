package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_GetRoleList(t *testing.T) {
	resp, _ := CreateClient().GetRoleList(-1, 0)
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_GetUsersInRole(t *testing.T) {
	roles, _ := CreateClient().GetRoleList(-1, 0)

	resp, _ := CreateClient().GetUsersInRole(roles.Result.RoleList[0].Roles[0].Id, -1, 0)
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_GetRoleGroup(t *testing.T) {
	roles, _ := CreateClient().GetRoleList(-1, 0)

	resp, _ := CreateClient().GetRoleGroup(roles.Result.RoleList[0].GroupId)
	t.Log(json.ToJson(resp))
}

func TestDingTalkClient_GetRoleDetail(t *testing.T) {
	roles, _ := CreateClient().GetRoleList(-1, 0)

	resp, _ := CreateClient().GetRoleDetail(roles.Result.RoleList[0].Roles[0].Id)
	t.Log(json.ToJson(resp))
}
