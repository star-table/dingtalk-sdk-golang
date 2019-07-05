package dingding_sdk_golang

type WorkNoticeRecallResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WorkNoticeMsg struct {
	MsgType    string            `json:"msgtype"`
	Text       *TextNotice       `json:"text"`
	Image      *ImageNotice      `json:"image"`
	Voice      *VoiceNotice      `json:"voice"`
	File       *FileNotice       `json:"file"`
	Link       *LinkNotice       `json:"link"`
	OA         *OANotice         `json:"oa"`
	Markdown   *MarkdownNotice   `json:"markdown"`
	ActionCard *ActionCardNotice `json:"action_card"`
}

type SendWorkNoticeResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	TaskId  int    `json:"task_id"`
}

type GetWorkNoticeProgressResp struct {
	ErrCode  int            `json:"errcode"`
	ErrMsg   string         `json:"errmsg"`
	Progress NoticeProgress `json:"progress"`
}

type GetWorkNoticeResultResp struct {
	ErrCode    int              `json:"errcode"`
	ErrMsg     string           `json:"errmsg"`
	SendResult NoticeSendResult `json:"send_result"`
}

type SendNormalNoticeResp struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	Receiver string `json:"receiver"`
}

type NoticeSendResult struct {
	InvalidUserIdList   []string `json:"invalid_user_id_list"`
	ForbiddenUserIdList []string `json:"forbidden_user_id_list"`
	FailedUserIdList    []string `json:"failed_user_id_list"`
	ReadUserIdList      []string `json:"read_user_id_list"`
	UnreadUserIdList    []string `json:"unread_user_id_list"`
	InvalidDeptIdList   []string `json:"invalid_dept_id_list"`
}

type NoticeProgress struct {
	ProgressInPercent int `json:"progress_in_percent"`
	Status            int `json:"status"`
}

type TextNotice struct {
	Content string `json:"content"`
}

type ImageNotice struct {
	MediaId string `json:"media_id"`
}

type VoiceNotice struct {
	MediaId  string `json:"media_id"`
	Duration string `json:"duration"`
}

type LinkNotice struct {
	MsgUrl string `json:"messageUrl"`
	PicUrl string `json:"picUrl"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type FileNotice struct {
	MediaId string `json:"media_id"`
}

type OANotice struct {
	MsgUrl string       `json:"message_url"`
	Head   OANoticeHead `json:"head"`
	Body   OANoticeBody `json:"body"`
}

type OANoticeHead struct {
	BgColor string `json:"bgcolor"`
	Text    string `json:"text"`
}

type OANoticeBody struct {
	Title     string             `json:"title"`
	Form      []OANoticeBodyForm `json:"form"`
	Rich      OANoticeBodyRich   `json:"rich"`
	Content   string             `json:"content"`
	Image     string             `json:"image"`
	FileCount string             `json:"file_count"`
	Author    string             `json:"author"`
}

type OANoticeBodyForm struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OANoticeBodyRich struct {
	Num  string `json:"num"`
	Unit string `json:"unit"`
}

type MarkdownNotice struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCardNotice struct {
	Title          string                   `json:"title"`
	Markdown       string                   `json:"markdown"`
	SingleTitle    string                   `json:"single_title"`
	SingleUrl      string                   `json:"single_url"`
	BtnOrientation int                      `json:"btn_orientation"`
	BtnJsonList    *[]ActionCardBtnJsonList `json:"btn_json_list"`
}

type ActionCardBtnJsonList struct {
	Title     string `json:"title"`
	ActionUrl string `json:"action_url"`
}
