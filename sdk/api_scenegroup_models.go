package sdk

type CreateSceneGroupReq struct {
	OwnerUserID         string `json:"owner_user_id"`
	UserIds             string `json:"user_ids"`
	UUID                string `json:"uuid"`
	Icon                string `json:"icon"`
	MentionAllAuthority int    `json:"mention_all_authority"`
	ShowHistoryType     int    `json:"show_history_type"`
	ValidationType      int    `json:"validation_type"`
	Searchable          int    `json:"searchable"`
	ChatBannedType      int    `json:"chat_banned_type"`
	ManagementType      int    `json:"management_type"`
	Title               string `json:"title"`
	TemplateID          string `json:"template_id"`
}

type CreateSceneGroupResp struct {
	BaseResp

	Result CreateSceneGroupResult `json:"result"`
}

type CreateSceneGroupResult struct {
	OpenConversationID string `json:"open_conversation_id"`
	ChatID             string `json:"chat_id"`
}
