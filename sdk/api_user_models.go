package sdk

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetDepMemberIdsResp struct {
	BaseResp
	UserIds []string `json:"userIds"`
}

type GetUserDetailResp struct {
	BaseResp

	UserList
}

type GetDepMemberListResp struct {
	BaseResp

	HasMore  bool       `json:"hasMore"`
	UserList []UserList `json:"userlist"`
}

type GetAdminListResp struct {
	BaseResp

	AdminList []AdminList `json:"admin_list"`
}

type GetAdminScopeResp struct {
	BaseResp

	DeptIds []int64 `json:"dept_ids"`
}

type CanAccessMicroAppResp struct {
	BaseResp

	CanAccess bool `json:"canAccess"`
}

type GetUserIdByUnionIdResp struct {
	BaseResp

	ContactType int    `json:"contactType"`
	UserId      string `json:"userid"`
}

type GetOrgUserCountResp struct {
	BaseResp

	Count int64 `json:"count"`
}

type AdminList struct {
	SysLevel int    `json:"sys_level"`
	UserId   string `json:"userid"`
}

type UserList struct {
	Name            string           `json:"name"`
	UnionId         string           `json:"unionid"`
	UserId          string           `json:"userid"`
	IsLeaderInDepts string           `json:"isLeaderInDepts"`
	IsBoos          bool             `json:"isBoss"`
	HiredDate       int64            `json:"hiredDate"`
	IsSenior        bool             `json:"isSenior"`
	Department      []int64          `json:"department"`
	OrderInDepts    string           `json:"orderInDepts"`
	Active          bool             `json:"active"`
	Avatar          string           `json:"avatar"`
	IsAdmin         bool             `json:"isAdmin"`
	IsHide          bool             `json:"isHide"`
	JobNumber       string           `json:"jobnumber"`
	Position        string           `json:"position"`
	DingId          string           `json:"dingId"`
	Roles           []UserDetailRole `json:"roles"`
}

type UserDetailRole struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
}

type GetDeptUserListV2Resp struct {
	BaseResp

	Result GetDeptUserListV2RespResult `json:"result"`
}

type GetDeptUserListV2RespResult struct {
	HasMore    bool               `json:"has_more"`
	NextCursor int64              `json:"nextCursor"`
	List       []UserDetailInfoV2 `json:"list"`
}

type UserDetailInfoV2 struct {
	UserID               string  `json:"userid"`
	UnionID              string  `json:"unionid"`
	Name                 string  `json:"name"`
	Avatar               string  `json:"avatar"`
	StateCode            string  `json:"state_code"`
	Mobile               string  `json:"mobile"`
	HideMobile           bool    `json:"hide_mobile"`
	Telephone            string  `json:"telephone"`
	JobNumber            string  `json:"job_number"`
	Title                string  `json:"title"`
	Email                string  `json:"email"`
	OrgEmail             string  `json:"org_email"`
	WorkPlace            string  `json:"work_place"`
	Remark               string  `json:"remark"`
	DeptIdList           []int64 `json:"dept_id_list"`
	DeptOrder            int64   `json:"dept_order"`
	Extension            string  `json:"extension"`
	HiredDate            int64   `json:"hired_date"`
	Active               bool    `json:"active"`
	Admin                bool    `json:"admin"`
	Boss                 bool    `json:"boss"`
	Leader               bool    `json:"leader"`
	ExclusiveAccount     bool    `json:"exclusive_account"`
	LoginId              string  `json:"login_id"`
	ExclusiveAccountType string  `json:"exclusive_account_type"`
}
