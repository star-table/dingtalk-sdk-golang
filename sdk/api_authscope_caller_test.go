package sdk

import (
	"testing"

	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

func TestGetAuthScopes(t *testing.T) {
	resp, err := CreateClient().GetAuthScopes()
	t.Log(json.ToJson(resp))
	t.Log(err)
}
