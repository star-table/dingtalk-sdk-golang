package sdk

type GetAuthScopesResp struct {
	BaseResp
	AuthUserField  []string `json:"auth_user_field"`
	ConditionField []string `json:"condition_field"`
	AuthOrgScopes  struct {
		AuthedDept []int64  `json:"authed_dept"`
		AuthedUser []string `json:"authed_user"`
	} `json:"auth_org_scopes"`
}
