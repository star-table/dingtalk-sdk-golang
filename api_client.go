package dingding_sdk_golang

import (
	"os"
)

type Corp struct {
	CorpId      string
	SuiteTicket string
	SuiteKey    string
	SuiteSecret string
}

type DingTalkClient struct {
	AccessToken string
	AgentId     int
}

func NewCorp(suiteTicket string, corpId string) *Corp {
	return &Corp{
		CorpId:      corpId,
		SuiteTicket: suiteTicket,
		SuiteKey:    os.Getenv("SUITE_KEY"),
		SuiteSecret: os.Getenv("SUITE_SECRET"),
	}
}

func NewDingTalkClient(accessToken string, agentId int) *DingTalkClient {
	return &DingTalkClient{
		AccessToken: accessToken,
		AgentId:     agentId,
	}

}

func (corp *Corp) CreateDingTalkClient() (*DingTalkClient, error) {
	tokenInfo, err := corp.GetCorpToken()
	if err != nil {
		return nil, err
	}
	authInfo, err := corp.GetAuthInfo()
	if err != nil {
		return nil, err
	}
	return NewDingTalkClient(tokenInfo.AccessToken, authInfo.AuthInfo.Agent[0].AgentId), nil
}

//Create Corp just for test
func CreateCorp() *Corp {
	os.Setenv("SUITE_KEY", "xxx")
	os.Setenv("SUITE_SECRET", "xxx")
	return NewCorp("xxx", "xxx")
}

//Create Client just for test
func CreateClient() *DingTalkClient {
	client, _ := CreateCorp().CreateDingTalkClient()
	return client
}
