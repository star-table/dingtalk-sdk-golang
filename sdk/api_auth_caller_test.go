package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/json"
	"os"
	"testing"
	"time"
)

//Create Corp just for test
func CreateCorp() *Corp {
	os.Setenv("SUITE_KEY", "suiteocpiljyoalvbhrbi")
	os.Setenv("SUITE_SECRET", "d1XKtyVpocDrOVJrDqPfqysmLGX7pinWS7iA8l5T7OWPd8aWZWNRfXEJrHoyb5Ng")
	return NewCorp("xxx", "ding79b4e083c47d2808f2c783f7214b6d69")
}

//Create Client just for test
func CreateClient() *DingTalkClient {
	client, _ := CreateCorp().CreateDingTalkClient()
	return client
}

func TestCorp_GetCorpToken(t *testing.T) {
	corpTokenInfo, _ := CreateCorp().GetCorpToken()
	t.Logf(json.ToJson(corpTokenInfo))
}

func TestCorp_GetAuthInfo(t *testing.T) {
	corpAuthInfo, _ := CreateCorp().GetAuthInfo()
	t.Logf(json.ToJson(corpAuthInfo))
}

func TestCorp_GetAgent(t *testing.T) {
	cp := CreateCorp()
	corpAuthInfo, _ := cp.GetAuthInfo()

	agentInfo, _ := cp.GetAgent(corpAuthInfo.AuthInfo.Agent[0].AgentId)
	t.Logf(json.ToJson(agentInfo))
}

func TestDingTalkClient_GetJsApiTicket(t *testing.T) {
	apiTicketInfo, err := CreateClient().GetJsApiTicket("1")
	t.Log(json.ToJson(apiTicketInfo))
	t.Log(err)
}

func TestCalculateJsApiSign(t *testing.T) {
	apiTicketInfo, err := CreateClient().GetJsApiTicket("1")
	if err != nil {
		t.Log(err)
	}
	ticket := apiTicketInfo.Ticket
	nonceStr := "abcdefg"
	timestamp := time.Now().UnixNano() / 1e6
	url := "http://192.168.1.126:8080/"
	signature := CalculateJsApiSign(ticket, nonceStr, timestamp, url)

	t.Logf("ticket: %s \n nonceStr: %s \n timestamp: %d \n url: %s \n signature: %s \n", ticket, nonceStr, timestamp, url, signature)
}
