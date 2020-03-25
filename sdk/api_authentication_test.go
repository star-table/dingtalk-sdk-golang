package sdk

import (
	"testing"

	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

func TestGetUserInfoFromThird(t *testing.T) {
	resp, err := CreateClient().GetUserInfoFromThird("f7e89bd957463f41a99882cddb55a34a")
	t.Log(json.ToJson(resp))
	t.Log(err)
}
