package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestGetCorpToken(t *testing.T) {
	corpTokenInfo, _ := CreateCorp().GetCorpToken()
	t.Logf(json.ToJson(corpTokenInfo))
}

func TestGetAuthInfo(t *testing.T) {
	corpAuthInfo, _ := CreateCorp().GetAuthInfo()
	t.Logf(json.ToJson(corpAuthInfo))
}

func TestGetAgent(t *testing.T) {
	cp := CreateCorp()
	corpAuthInfo, _ := cp.GetAuthInfo()

	agentInfo, _ := cp.GetAgent(corpAuthInfo.AuthInfo.Agent[0].AgentId)
	t.Logf(json.ToJson(agentInfo))
}
