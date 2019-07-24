package sdk

type SaveProcessResp struct {
	BaseResp
	Result struct {
		ProcessCode string `json:"process_code"`
	} `json:"result"`
}

type SaveProcessReq struct {
	ProcessCode       *string           `json:"process_code"`
	AgentId           int64             `json:"agentid"`
	DisableFormEdit   *string           `json:"disable_form_edit"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	FormComponentList []FormComponentVo `json:"form_component_list"`
}

type FormComponentVo struct {
	ComponentName string              `json:"component_name"`
	Props         FormComponentPropVo `json:"props"`
	Children      *FormComponentVo    `json:"children"`
}

type FormComponentPropVo struct {
	Id          string                 `json:"id"`
	Label       string                 `json:"label"`
	Required    *bool                  `json:"required"`
	NotPrint    *string                `json:"not_print"`
	Placeholder *string                `json:"placeholder"`
	NotUpper    *string                `json:"not_upper"`
	Unit        *string                `json:"unit"`
	Options     *[]string              `json:"options"`
	Choice      *int64                 `json:"choice"`
	Link        *string                `json:"link"`
	ActionName  *string                `json:"action_name"`
	StatField   *[]FormComponentStatVo `json:"stat_field"`
}

type FormComponentStatVo struct {
	Id    *string `json:"id"`
	Label *string `json:"label"`
	Upper *bool   `json:"upper"`
	Unit  *string `json:"unit"`
}

type FormConditionListResp struct {
	BaseResp
	Result struct {
		List struct {
			Id    string `json:"id"`
			Label string `json:"label"`
		} `json:"list"`
	} `json:"result"`
}

type CreateProcessInstanceReq struct {
	AgentId             *int64                       `json:"agent_id"`
	ProcessCode         string                       `json:"process_code"`
	OriginatorUserId    string                       `json:"originator_user_id"`
	DeptId              int64                        `json:"dept_id"`
	Approvers           string                       `json:"approvers"`
	ApproversV2         *[]ProcessInstanceApproverVo `json:"approvers_v2"`
	CcList              *string                      `json:"cc_list"`
	CcPosition          *string                      `json:"cc_position"`
	FormComponentValues []FormComponentValuesVo      `json:"form_component_values"`
}

type ProcessInstanceApproverVo struct {
	UserIds        *[]string `json:"user_ids"`
	TaskActionType *string   `json:"task_action_type"`
}

type FormComponentValuesVo struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value"`
}

type CreateProcessInstanceResp struct {
	BaseResp
	ProcessInstanceId string `json:"process_instance_id"`
}

type GetCspaceInfoResp struct {
	BaseResp
	Result struct {
		SpaceId int64 `json:"space_id"`
	} `json:"result"`
	Success bool `json:"success"`
}

type GetProcessInstanceInfoResp struct {
	BaseResp
	ProcessInstance ProcessInstanceVo `json:"process_instance"`
}

type ProcessInstanceVo struct {
	Title                      string                  `json:"title"`
	CreateTime                 string                  `json:"create_time"`
	FinishTime                 string                  `json:"finish_time"`
	OriginatorUserId           string                  `json:"originator_userid"`
	OriginatorDeptId           string                  `json:"originator_dept_id"`
	Status                     string                  `json:"status"`
	CcUserIds                  string                  `json:"cc_userids"`
	FormComponentValues        []FormComponentValuesVo `json:"form_component_values"`
	Result                     string                  `json:"result"`
	BusinessId                 string                  `json:"business_id"`
	OperationRecords           []OperationRecordsVo    `json:"operation_records"`
	Tasks                      []TasksVo               `json:"tasks"`
	OriginatorDeptName         string                  `json:"originator_dept_name"`
	BizAction                  string                  `json:"biz_action"`
	AttachedProcessInstanceIds []string                `json:"attached_process_instance_ids"`
}

type OperationRecordsVo struct {
	UserId          string `json:"userid"`
	Date            string `json:"date"`
	OperationResult string `json:"operation_result"`
	Remark          string `json:"remark"`
}

type TasksVo struct {
	UserId     string `json:"userid"`
	TaskStatus string `json:"task_status"`
	TaskResult string `json:"task_result"`
	CreateTime string `json:"create_time"`
	FinishTime string `json:"finish_time"`
	TaskId     string `json:"taskid"`
}
