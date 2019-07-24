package sdk

//======================== 获取企业凭证: https://open-doc.dingtalk.com/microapp/serverapi3/vfitg0
type GetCorpTokenResp struct {
	BaseResp

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//======================== 获取企业授权信息: https://open-doc.dingtalk.com/microapp/serverapi3/hv357q
type GetAuthInfoResp struct {
	BaseResp

	AuthCorpInfo    AuthCorpInfo    `json:"auth_corp_info"`
	AuthUserInfo    AuthUserInfo    `json:"auth_user_info"`
	AuthInfo        AuthInfo        `json:"auth_info"`
	ChannelAuthInfo ChannelAuthInfo `json:"channel_auth_info"`
}

type AuthCorpInfo struct {
	CorpLogoUrl     string `json:"corp_logo_url"`
	CorpName        string `json:"corp_name"`
	CorpId          string `json:"corpid"`
	Industry        string `json:"industry"`
	InviteCode      string `json:"invite_code"`
	LicenseCode     string `json:"license_code"`
	AuthChannel     string `json:"auth_channel"`
	AuthChannelType string `json:"auth_channel_type"`
	IsAuthenticated bool   `json:"is_authenticated"`
	AuthLevel       int64  `json:"auth_level"`
	InviteUrl       string `json:"invite_url"`
	CorpProvince    string `json:"corp_province"`
	CorpCity        string `json:"corp_city"`
}

type AuthUserInfo struct {
	UserId string `json:"userId"`
}

type AuthInfo struct {
	Agent []Agent `json:"agent"`
}

type Agent struct {
	AgentName string   `json:"agent_name"`
	AgentId   int64    `json:"agentid"`
	AppId     int64    `json:"appid"`
	LogoUrl   string   `json:"logo_url"`
	AdminList []string `json:"admin_list"`
}

type ChannelAuthInfo struct {
	ChannelAgent []ChannelAgent `json:"channelAgent"`
}

type ChannelAgent struct {
	AgentName string `json:"agent_name"`
	AgentId   int64  `json:"agentid"`
	AppId     int64  `json:"appid"`
	LogoUrl   int64  `json:"logo_url"`
}

//======================== 获取授权应用信息: https://open-doc.dingtalk.com/microapp/serverapi3/vfitg0
type GetAgentResp struct {
	BaseResp

	AgentId     int64  `json:"agentid"`
	Name        string `json:"name"`
	LogoUrl     string `json:"logo_url"`
	Description string `json:"description"`
	Close       int64  `json:"close"`
}

type GetJsApiTicketResp struct {
	BaseResp

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

type GetSuiteAccessTokenResp struct {
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int64  `json:"expires_in"`
}

type GetPermanentCodeResp struct {
	PermanentCode string       `json:"permanent_code"`
	AuthCorpInfo  AuthCorpInfo `json:"auth_corp_info"`
}
