package dingding_sdk_golang

type GetUserInfoFromThirdResp struct {
	BaseResp
	UserId   string `json:"userid"`
	SysLevel int64  `json:"sys_level"`
	IsSys    bool   `json:"is_sys"`
	DeviceId string `json:"deviceId"`
}

type GetUserInfoFromAdminResp struct {
	BaseResp
	CropInfo struct {
		CropName string `json:"crop_name"`
		CropId   string `json:"cropid"`
	}
	IsSys    bool `json:"is_sys"`
	UserInfo struct {
		Avatar string `json:"avatar"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		UserId string `json:"userid"`
	} `json:"user_info"`
}

type GetUserInfoByCodeResp struct {
	BaseResp
	UserInfo struct {
		Nick    string `json:"nick"`
		OpenId  string `json:"openid"`
		UnionId string `json:"unionid"`
	} `json:"user_info"`
}
