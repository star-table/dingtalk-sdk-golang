package sdk

type UploadMediaResp struct {
	BaseResp

	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}
