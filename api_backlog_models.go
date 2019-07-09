package dingding_sdk_golang

type CreateOrUpdateBackLogResp struct {
	BaseResp

	Result CreateOrUpdateBackLogResult `json:"result"`
}

type CreateWorkRecordResp struct {
	BaseResp

	Result CreateWorkRecordResult `json:"result"`
}

type CreateWorkRecordResult struct {
	ProcessInstanceId string `json:"process_instance_id"`
}

type CreateOrUpdateBackLogResult struct {
	ProcessCode string `json:"process_code"`
}

type CreateOrUpdateBackLogReq struct {
	SaveProcessRequest SaveProcessRequest `json:"SaveProcessRequest"`
}

type DeleteBackLogReq struct {
	AgentId     int    `json:"agentid"`
	ProcessCode string `json:"process_code"`
}

type CreateWorkRecordRequest struct {
	AgentId             int                   `json:"agentid"`
	ProcessCode         string                `json:"process_code"`
	OriginatorUserId    string                `json:"originator_user_id"`
	FormComponentValues []FormComponentValues `json:"form_component_values"`
	Url                 string                `json:"url"`
}

type UpdateWorkRecordRequest struct {
	AgentId           int    `json:"agentid"`
	ProcessInstanceId string `json:"process_instance_id"`
	Status            string `json:"status"`
	Result            string `json:"result,omitempty"`
}

type FormComponentValues struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SaveProcessRequest struct {
	AgentId           int                 `json:"agentid"`
	ProcessCode       string              `json:"process_code,omitempty"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	FormComponentList []FormComponentList `json:"form_component_list"`
	FakeMode          bool                `json:"fake_mode"`
}

type FormComponentList struct {
	ComponentName string `json:"component_name"`
	Props         Props  `json:"props"`
}

type Props struct {
	Id          string `json:"id"`
	Label       string `json:"label"`
	Required    bool   `json:"required"`
	Unit        string `json:"unit,omitempty"`
	Placeholder string `json:"placeholder"`
}
