package sdk

type GetSubdeptResp struct {
	BaseResp
	SubDeptIdList []int `json:"sub_dept_id_list"`
}

type GetDeptListResp struct {
	BaseResp
	Department []struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		ParentId        int    `json:"parentid"`
		CreateDeptGroup bool   `json:"createDeptGroup"`
		AutoAddUser     bool   `json:"autoAddUser"`
	} `json:"department"`
}

type GetDeptDetailResp struct {
	BaseResp
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Order                 int64  `json:"order"`
	ParentId              int    `json:"parentid"`
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
	ParentIds []int `json:"parentIds"`
}

type ListParentDeptsResp struct {
	BaseResp
	Department [][]int `json:"department"`
}
