package sdk

type GetAuthScopesResp struct {
	BaseResp
	AuthUserField  []string          `json:"auth_user_field"`
	ConditionField []string          `json:"condition_field"`
	AuthOrgScopes  GetAuthScopesData `json:"auth_org_scopes"`
}

type GetAuthScopesData struct {
	AuthedDept []int64  `json:"authed_dept"`
	AuthedUser []string `json:"authed_user"`
}
