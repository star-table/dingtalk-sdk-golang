package sdk

type GetRoleListResp struct {
	BaseResp

	Result GetRoleListResult `json:"result"`
}

type GetUsersInRoleResp struct {
	BaseResp

	Result GetUserInRoleResult `json:"result"`
}

type GetRoleGroupResp struct {
	BaseResp

	RoleGroup RoleGroupV1 `json:"role_group"`
}

type GetRoleDetailResp struct {
	BaseResp

	Role Role `json:"role"`
}

type GetUserInRoleResult struct {
	HasMore    bool       `json:"hasMore"`
	NextCursor int        `json:"nextCursor"`
	UserList   []UserList `json:"list"`
}

type GetRoleListResult struct {
	HasMore  bool        `json:"hasMore"`
	RoleList []RoleGroup `json:"list"`
}

type RoleGroup struct {
	Name    string `json:"name"`
	GroupId int    `json:"groupId"`
	Roles   []Role `json:"roles"`
}

type RoleGroupV1 struct {
	Name  string   `json:"group_name"`
	Roles []RoleV1 `json:"roles"`
}

type RoleV1 struct {
	Name string `json:"role_name"`
	Id   int    `json:"role_id"`
}

type Role struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
