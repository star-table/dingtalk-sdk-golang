package sdk

type ListLabelGroupsResp struct {
	BaseResp
	Results []struct {
		Name   string   `json:"name"`
		Color  int64    `json:"color"`
		Labels []Labels `json:"labels"`
	}
}

type Labels struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

type GetExtcontactListResp struct {
	BaseResp
	Results []ExtcontactDetail `json:"results"`
}

type GetExtcontactDetailResp struct {
	BaseResp
	Result ExtcontactDetail `json:"result"`
}

type ExtcontactDetail struct {
	Title          string   `json:"title"`
	ShareDeptIds   []int64  `json:"share_dept_ids"`
	LabelIds       []int64  `json:"label_ids"`
	Remark         string   `json:"remark"`
	Address        string   `json:"address"`
	Name           string   `json:"name"`
	FollowerUserId string   `json:"follower_user_id"`
	StateCode      string   `json:"state_code"`
	CompanyName    string   `json:"company_name"`
	ShareUserIds   []string `json:"share_user_ids"`
	UserId         string   `json:"userid"`
}
