package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestGetDepMember(t *testing.T) {

	corpTokenInfo, _ := GetCorpToken(ACCESS_KEY, SUITE_SECRET, SUITE_TICKET, AUTH_CORP_ID)
	accessToken := corpTokenInfo.AccessToken

	memberList, _ := GetDepMember(accessToken, "1")
	t.Logf(json.ToJson(memberList))
}
