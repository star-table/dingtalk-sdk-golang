package dingding_sdk_golang

type BaseResp struct {
	ErrCode int    `json:"errcode""`
	ErrMsg  string `json:"errmsg"`
}

type GetSubdeptResp struct {
	BaseResp
	SubDeptIdList []string `json:"subDeptIdList"`
}

type GetDeptListResp struct {
	BaseResp
	Department []struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		ParentId        string `json:"parentId"`
		CreateDeptGroup bool   `json:"createDeptGroup"`
		AutoAddUser     bool   `json:"autoAddUser"`
	}
}

type GetDeptDetailResp struct {
	BaseResp
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Order                int64  `json:"order"`
	CreateDeptGroup      bool   `json:"createDeptGroup"`
	AutoAddUser          bool   `json:"autoAddUser"`
	DeptHiding           bool   `json:"deptHiding"`
	DeptPermits          string `json:"deptPermits"`
	UserPermits          string `json:"userPermits"`
	OuterDept            bool   `json:"outerDept"`
	OuterPermitDepts     string `json:"outerPermitDepts"`
	OuterPermitUsers     string `json:"outerPermitUsers"`
	OrgDeptOwner         string `json:"orgDeptOwner"`
	DeptManageUseridList string `json:"deptManageUseridList"`
	SourceIdentifier     string `json:"sourcIdentifier"`
}

type ListParentDeptsByDeptResp struct {
	BaseResp
	ParentIds []string `json:"parentIds"`
}

type ListParentDeptsResp struct {
	BaseResp
	Department [][]string
}
