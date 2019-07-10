package dingding_sdk_golang

type RegisterCallBackReq struct {
	CallBackTag []string `json:"call_back_tag"`
	Token       string   `json:"token"`
	AesKey      string   `json:"aes_key"`
	Url         string   `json:"url"`
}
