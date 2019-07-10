package sdk

import (
	"github.com/polaris-team/dingding-sdk-golang/json"
	"testing"
)

func TestSaveProcess(t *testing.T) {
	client := CreateClient()
	saveProcessrequest := SaveProcessReq{
		AgentId:     client.AgentId,
		Name:        "人工测试4",
		Description: "描述信息",
		FormComponentList: []FormComponentVo{
			{
				ComponentName: "TextField",
				Props: FormComponentPropVo{
					Id:    "TextField-J78F056R",
					Label: "第一栏",
				},
			},
		},
	}
	resp, err := client.SaveProcess(saveProcessrequest, nil, nil)
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestFormConditionList(t *testing.T) {
	resp, err := CreateClient().FormConditionList("PROC-848D9EE1-AD41-46F9-88CA-8AB33693F61D", nil)
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestCreateProcessInstance(t *testing.T) {
	memberList, _ := CreateClient().GetDepMemberIds("1")
	t.Log(json.ToJson(memberList))
	createProcessInstanceRequest := CreateProcessInstanceReq{
		ProcessCode:      "PROC-848D9EE1-AD41-46F9-88CA-8AB33693F61D",
		OriginatorUserId: memberList.UserIds[1],
		DeptId:           1,
		FormComponentValues: []FormComponentValuesVo{
			{
				Name:  "第一栏",
				Value: "what?",
			},
		},
	}

	resp, err := CreateClient().CreateProcessInstance(createProcessInstanceRequest)
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetProcessInstanceInfo(t *testing.T) {
	resp, err := CreateClient().GetProcessInstanceInfo("7d335a66-4d66-4c98-b4da-94949ebebcaf")
	t.Log(json.ToJson(resp))
	t.Log(err)
}

func TestGetCspaceInfo(t *testing.T) {
	memberList, _ := CreateClient().GetDepMemberIds("1")
	t.Log(json.ToJson(memberList))
	resp, err := CreateClient().GetCspaceInfo(memberList.UserIds[1])
	t.Log(json.ToJson(resp))
	t.Log(err)
}
