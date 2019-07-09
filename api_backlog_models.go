package dingding_sdk_golang

type CreateOrUpdateBackLogResp struct {
	BaseResp

	Result CreateOrUpdateBackLogResult `json:"result"`
}

type CreateOrUpdateBackLogResult struct {
	ProcessCode string `json:"process_code"`
}

type CreateOrUpdateBackLogReq struct {
	SaveProcessRequest SaveProcessRequest `json:"SaveProcessRequest"`
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
