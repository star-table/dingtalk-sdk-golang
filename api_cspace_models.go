package dingding_sdk_golang

type FileUploadSingleResp struct {
	BaseResp

	MediaId string `json:"media_id"`
}

type CommitUploadTransaction struct {
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

type BeginUploadTransaction struct {
	BaseResp
	UploadId string `json:"upload_id"`
}
