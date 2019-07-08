package dingding_sdk_golang

import "testing"

func TestGetSubDept(t *testing.T) {
	resp, err := CreateClient().GetSubDept("1")
	t.Log(resp, err)
}

func TestGetDeptList(t *testing.T) {
	fetchChild := false
	resp, err := CreateClient().GetDeptList(nil, &fetchChild, 1)
	t.Log(resp, err)
}

func TestGetDeptDetail(t *testing.T) {
	resp, err := CreateClient().GetDeptDetail("1", nil)
	t.Log(resp, err)
}

func TestListParentDeptsByDept(t *testing.T) {
	resp, err := CreateClient().ListParentDeptsByDept("1")
	t.Log(resp, err)
}

func TestListParentDepts(t *testing.T) {
	resp, err := CreateClient().ListParentDepts("1")
	t.Log(resp, err)
}
