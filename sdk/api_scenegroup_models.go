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

type UpdateSceneGroupReq struct {
	OpenConversationID string `json:"open_conversation_id"`

	OwnerUserID         *string `json:"owner_user_id,omitempty"`
	UserIds             *string `json:"user_ids,omitempty"`
	UUID                *string `json:"uuid,omitempty"`
	Icon                *string `json:"icon,omitempty"`
	MentionAllAuthority *int    `json:"mention_all_authority,omitempty"`
	ShowHistoryType     *int    `json:"show_history_type,omitempty"`
	ValidationType      *int    `json:"validation_type,omitempty"`
	Searchable          *int    `json:"searchable,omitempty"`
	ChatBannedType      *int    `json:"chat_banned_type,omitempty"`
	ManagementType      *int    `json:"management_type,omitempty"`
	Title               *string `json:"title,omitempty"`
}

type GetSceneGroupReq struct {
	OpenConversationID string `json:"open_conversation_id"`
}

type GetSceneGroupResp struct {
	BaseResp

	Result GetSceneGroupResult `json:"result"`
}

type GetSceneGroupResult struct {
	Icon               string            `json:"icon"`
	ManagementOptions  ManagementOptions `json:"management_options"`
	Title              string            `json:"title"`
	TemplateID         string            `json:"template_id"`
	OpenConversationID string            `json:"open_conversation_id"`
	SubAdminStaffIds   []string          `json:"sub_admin_staff_ids"`
	OwnerStaffID       string            `json:"owner_staff_id"`
	GroupUrl           string            `json:"group_url"`
}

type ManagementOptions struct {
	ChatBannedType      string `json:"chat_banned_type"`
	Searchable          string `json:"searchable"`
	ValidationType      string `json:"validation_type"`
	MentionAllAuthority string `json:"mention_all_authority"`
	ManagementType      string `json:"management_type"`
	ShowHistoryType     string `json:"show_history_type"`
}

type AddSceneGroupMemberReq struct {
	OpenConversationID string `json:"open_conversation_id"`
	UserIds            string `json:"user_ids"`
}

type DeleteSceneGroupMemberReq struct {
	OpenConversationID string `json:"open_conversation_id"`
	UserIds            string `json:"user_ids"`
}
