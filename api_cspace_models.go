package dingding_sdk_golang

type FileUploadSingleResp struct {
	BaseResp

	MediaId string `json:"media_id"`
}

type GetCustomSpaceResp struct {
	BaseResp
	SpaceId string `json:"spaceid"`
}

type AddFileToUserCSpaceResp struct {
	BaseResp
	DEntry string `json:"dentry"`
}
