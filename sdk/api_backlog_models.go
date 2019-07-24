package sdk

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
	AgentId     int64  `json:"agentid"`
	ProcessCode string `json:"process_code"`
}

type CreateWorkRecordRequest struct {
	AgentId             int64                 `json:"agentid"`
	ProcessCode         string                `json:"process_code"`
	OriginatorUserId    string                `json:"originator_user_id"`
	FormComponentValues []FormComponentValues `json:"form_component_values"`
	Url                 string                `json:"url"`
}

type UpdateWorkRecordRequest struct {
	AgentId           int64  `json:"agentid"`
	ProcessInstanceId string `json:"process_instance_id"`
	Status            string `json:"status"`
	Result            string `json:"result,omitempty"`
}

type CreateWorkRecordTaskRequest struct {
	AgentId           int64                     `json:"agentid"`
	ProcessInstanceId string                    `json:"process_instance_id"`
	ActivityId        *string                   `json:"activity_id,omitempty"`
	Tasks             []CreateWorkRecordTaskTop `json:"tasks"`
}

type UpdateWorkRecordTaskRequest struct {
	AgentId           int64                     `json:"agentid"`
	ProcessInstanceId string                    `json:"process_instance_id"`
	Tasks             []UpdateWorkRecordTaskTop `json:"tasks"`
}

type CancelTaskGroupRequest struct {
	AgentId           int64   `json:"agentid"`
	ProcessInstanceId string  `json:"process_instance_id"`
	ActivityId        *string `json:"activity_id"`
}

type UpdateWorkRecordTaskTop struct {
	TaskId int64  `json:"task_id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}

type CreateWorkRecordTaskResp struct {
	BaseResp

	Tasks []CreateWorkRecordTaskRespTasks `json:"tasks"`
}

type UpdateWorkRecordTaskResp struct {
	CreateWorkRecordTaskResp
}

type CreateWorkRecordTaskRespTasks struct {
	TaskId int64  `json:"task_id"`
	UserId string `json:"userid"`
}

type CreateWorkRecordTaskTop struct {
	UserId string `json:"userid"`
	Url    string `json:"url"`
}

type FormComponentValues struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SaveProcessRequest struct {
	AgentId           int64               `json:"agentid"`
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
