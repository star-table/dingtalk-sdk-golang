package sdk

import (
	"github.com/polaris-team/dingtalk-sdk-golang/json"
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
