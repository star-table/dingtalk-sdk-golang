package sdk

import (
	"github.com/polaris-team/dingtalk-sdk-golang/json"
	"testing"
)

func TestGetSubDept(t *testing.T) {
	resp, err := CreateClient().GetSubDept("1")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetDeptList(t *testing.T) {
	fetchChild := false
	resp, err := CreateClient().GetDeptList(nil, &fetchChild, "1")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetDeptDetail(t *testing.T) {
	resp, err := CreateClient().GetDeptDetail("1", nil)
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestListParentDeptsByDept(t *testing.T) {
	resp, err := CreateClient().ListParentDeptsByDept("1")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestListParentDepts(t *testing.T) {
	memberList, _ := CreateClient().GetDepMemberIds("1")
	t.Log(json.ToJson(memberList))
	resp, err := CreateClient().ListParentDepts(memberList.UserIds[1])
	t.Log(json.ToJson(resp))
	t.Log(err)
}
