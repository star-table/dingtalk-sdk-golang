package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestGetCorpToken(t *testing.T) {
	corpTokenInfo, _ := GetCorpToken(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	t.Logf(json.ToJson(corpTokenInfo))
}

func TestGetAuthInfo(t *testing.T) {
	corpAuthInfo, _ := GetAuthInfo(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	t.Logf(json.ToJson(corpAuthInfo))
}

func TestGetAgent(t *testing.T) {
	agentInfo, _ := GetAgent(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID, 273673341)
	t.Logf(json.ToJson(agentInfo))
}
