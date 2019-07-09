package dingding_sdk_golang

type GetUserInfoResp struct {
	BaseResp
	UserId string `json:"userid"`
	SysLevel int64 `json:"sys_level"`
	IsSys bool `json:"is_sys"`
	DeviceId string `json:"deviceId"`
}