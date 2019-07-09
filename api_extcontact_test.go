package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestListLabelGroups(t *testing.T) {
	resp, err := CreateClient().ListLabelGroups("20", "0")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetExtcontactList(t *testing.T) {
	resp, err := CreateClient().GetExtcontactList("20", "0")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetExtcontactDetail(t *testing.T) {
	userInfo, _ := CreateClient().GetExtcontactList("20", "0")
	t.Log(json.ToJson(userInfo))
	resp, err := CreateClient().GetExtcontactDetail(userInfo.Results[0].UserId)
	t.Log(json.ToJson(resp))
	t.Log(err)
}