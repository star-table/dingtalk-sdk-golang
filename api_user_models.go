package dingding_sdk_golang

type GetDepMemberResp struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	UserIds []string `json:"userIds"`
}

type GetUserDetailResp struct {
	ErrCode         int    `json:"errcode"`
	ErrMsg          string `json:"errmsg"`
	Name            string ``
	UnionId         string
	UserId          string
	IsLeaderInDepts string
	IsBoos          bool
	HiredDate       int64
	IsSenior        bool
	Department      []int
}
