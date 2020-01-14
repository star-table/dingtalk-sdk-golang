package sdk

type GetSubdeptResp struct {
	BaseResp
	SubDeptIdList []int64 `json:"sub_dept_id_list"`
}

type GetDeptListResp struct {
	BaseResp
	Department []DepartmentInfo `json:"department"`
}

type DepartmentInfo struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	ParentId        int64  `json:"parentid"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	AutoAddUser     bool   `json:"autoAddUser"`
}

type GetDeptDetailResp struct {
	BaseResp
	Id                    int64  `json:"id"`
	Name                  string `json:"name"`
	Order                 int64  `json:"order"`
	ParentId              int64  `json:"parentid"`
	CreateDeptGroup       bool   `json:"createDeptGroup"`
	AutoAddUser           bool   `json:"autoAddUser"`
	DeptHiding            bool   `json:"deptHiding"`
	DeptPermits           string `json:"deptPermits"`
	UserPermits           string `json:"userPermits"`
	OuterDept             bool   `json:"outerDept"`
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	OrgDeptOwner          string `json:"orgDeptOwner"`
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	SourceIdentifier      string `json:"sourceIdentifier"`
}

type ListParentDeptsByDeptResp struct {
	BaseResp
	ParentIds []int64 `json:"parentIds"`
}

type ListParentDeptsResp struct {
	BaseResp
	Department [][]int64 `json:"department"`
}
