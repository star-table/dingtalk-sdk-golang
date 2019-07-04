package dingding_sdk_golang

type GetDepMemberResp struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	UserIds []string `json:"userIds"`
}
