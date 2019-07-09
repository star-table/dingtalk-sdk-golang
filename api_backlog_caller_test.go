package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestDingTalkClient_CreateOrUpdateBackLog(t *testing.T) {
	client := CreateClient()

	req := SaveProcessRequest{
		Name:        "您的垃圾需要丢掉了，真的5~",
		Description: "垃圾桶已满，请及时丢掉.",
		FakeMode:    true,
		FormComponentList: []FormComponentList{
			{
				ComponentName: "TextareaField",
				Props: Props{
					Id:          "TextareaField-J78F077S",
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

	req.ProcessCode = resp.Result.ProcessCode
	req.Description += "test"

	resp, err = client.CreateOrUpdateBackLog(req)
	respJson, _ := json.ToJson(resp)
	t.Logf("update: %s", respJson)
	t.Log(err)
}

func TestDingTalkClient_DeleteBackLog(t *testing.T) {
	client := CreateClient()
	req := DeleteBackLogReq{
		ProcessCode: "PROC-98F37BF9-3724-49FB-BE3B-E4321A5451D1",
	}
	resp, err := client.DeleteBackLog(req)
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestDingTalkClient_CreateAndUpdateWorkRecord(t *testing.T) {
	client := CreateClient()

	req := CreateWorkRecordRequest{
		ProcessCode:      "PROC-9EDD9794-2C7F-458D-ABAF-394D492CAABB",
		OriginatorUserId: "15001956402427783",
		FormComponentValues: []FormComponentValues{
			{
				Name:  "username",
				Value: "nico",
			},
		},
		Url: "http://study.ikuvn.com",
	}
	title := "倒垃圾待办实例"
	resp, err := client.CreateWorkRecord(req, &title)
	t.Log(json.ToJson(resp))
	t.Log(err)

	updateReq := UpdateWorkRecordRequest{
		ProcessInstanceId: resp.Result.ProcessInstanceId,
		Status:            "COMPLETED",
		Result:            "agree",
	}
	updateResp, err := client.UpdateWorkRecord(updateReq)
	t.Log(json.ToJson(updateResp))
	t.Log(err)

	resp, err = client.CreateWorkRecord(req, &title)
	t.Log(json.ToJson(resp))
	t.Log(err)
	updateReq.ProcessInstanceId = resp.Result.ProcessInstanceId

	updateReq.Result = "refuse"
	updateResp, err = client.UpdateWorkRecord(updateReq)
	t.Log(json.ToJson(updateResp))
	t.Log(err)

	resp, err = client.CreateWorkRecord(req, &title)
	t.Log(json.ToJson(resp))
	t.Log(err)
	updateReq.ProcessInstanceId = resp.Result.ProcessInstanceId

	updateReq.Status = "TERMINATED"
	updateResp, err = client.UpdateWorkRecord(updateReq)
	t.Log(json.ToJson(updateResp))
	t.Log(err)
}
