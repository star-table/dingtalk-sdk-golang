package dingding_sdk_golang

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"reflect"
	"strconv"
)

//创建或更新审批模板
//https://open-doc.dingtalk.com/microapp/serverapi3/svyi4c
func (client *DingTalkClient) SaveProcess(saveProcessRequest SaveProcessReq, fakeMode *bool, templateEditUrl *string) (SaveProcessResp, error) {
	params := make(map[string]string)
	saveProcessRequest.AgentId = client.AgentId
	saveProcess, _ := json.ToJson(saveProcessRequest)
	params["saveProcessRequest"] = encrypt.URLEncode(saveProcess)
	params["access_token"] = client.AccessToken
	if fakeMode != nil {
		params["fake_mode"] = strconv.FormatBool(*fakeMode)
	}
	if templateEditUrl != nil {
		params["template_edit_url"] = *templateEditUrl
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/process/save", params, "")
	resq := SaveProcessResp{}
	if err != nil {
		return resq, err
	}
	json.FromJson(body, &resq)
	return resq, err
}

//查询已设置为条件的表单组件
//https://open-doc.dingtalk.com/microapp/serverapi3/yqcw6y
func (client *DingTalkClient) FormConditionList(processCode string, agentId *int) (FormConditionListResp, error) {
	request := map[string]string{
		"process_code": processCode,
	}
	if agentId != nil {
		request["agentid"] = strconv.Itoa(*agentId)
	}
	requestJson, _ := json.ToJson(request)
	params := map[string]string{
		"access_token": client.AccessToken,
		"request":      encrypt.URLEncode(requestJson),
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/process/form/condition/list", params, "")
	resq := FormConditionListResp{}
	if err != nil {
		return resq, err
	}
	json.FromJson(body, &resq)
	return resq, err
}

//发起审批实例
//https://open-doc.dingtalk.com/microapp/serverapi3/zzerip
func (client *DingTalkClient) CreateProcessInstance(createParams CreateProcessInstanceReq) (CreateProcessInstanceResp, error) {
	params := make(map[string]string)
	typ := reflect.TypeOf(createParams)
	val := reflect.ValueOf(createParams)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		v := ""
		tname := field.Type.Name()
		if value != nil {
			if tname == "string" {
				v = value.(string)
			} else if tname == "int" {
				v = strconv.Itoa(value.(int))
			} else {
				if !reflect2.IsNil(value) {
					v, _ = json.ToJson(value)
					v = encrypt.URLEncode(v)
				}
			}
		}
		if v != "" {
			params[field.Tag.Get("json")] = v
		}
	}
	params["access_token"] = client.AccessToken
	if params["agent_id"] == "" {
		params["agent_id"] = strconv.Itoa(client.AgentId)
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/processinstance/create", params, "")
	resq := CreateProcessInstanceResp{}
	if err != nil {
		return resq, err
	}
	json.FromJson(body, &resq)
	return resq, err
}

//获取审批实例详情
//https://open-doc.dingtalk.com/microapp/serverapi3/tkgfol
func (client *DingTalkClient) GetProcessInstanceInfo(processInstanceId string) (GetProcessInstanceInfoResp, error) {
	params := map[string]string{
		"access_token":        client.AccessToken,
		"process_instance_id": processInstanceId,
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/processinstance/get", params, "")
	resq := GetProcessInstanceInfoResp{}
	if err != nil {
		return resq, err
	}
	json.FromJson(body, &resq)
	return resq, err
}

//获取审批钉盘空间信息
//https://open-doc.dingtalk.com/microapp/serverapi3/dnh74x
func (client *DingTalkClient) GetCspaceInfo(userId string) (GetCspaceInfoResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"user_id":      userId,
		"agent_id":     strconv.Itoa(client.AgentId),
	}

	body, err := http.Post("https://oapi.dingtalk.com/topapi/processinstance/cspace/info", params, "")
	fmt.Println(json.ToJson(body))
	resq := GetCspaceInfoResp{}
	if err != nil {
		return resq, err
	}
	json.FromJson(body, &resq)
	return resq, err
}
