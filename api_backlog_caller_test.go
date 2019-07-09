package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_CreateOrUpdateBackLog(t *testing.T) {
	client := CreateClient()

	req := SaveProcessRequest{
		AgentId:     client.AgentId,
		Name:        "您的垃圾需要丢掉了，真的~",
		Description: "垃圾桶已满，请及时丢掉.",
		FakeMode:    true,
		FormComponentList: []FormComponentList{
			{
				ComponentName: "TextareaField",
				Props: Props{
					Id:          "TextareaField-J78F057S",
					Required:    true,
					Label:       "username",
					Placeholder: "Please enter username",
				},
			},
		},
	}
	resp, err := client.CreateOrUpdateBackLog(req)
	t.Log(json.ToJson(resp))
	t.Log(err)
}
