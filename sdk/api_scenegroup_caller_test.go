package sdk

import (
	"github.com/polaris-team/dingtalk-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_UpdateSceneGroup(t *testing.T) {
	req := UpdateSceneGroupReq{}
	t.Log(json.ToJson(req))
}
