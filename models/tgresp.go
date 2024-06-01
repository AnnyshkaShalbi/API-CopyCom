package models

type TgResp_Updates struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id      int64 `json:"update_id"`
		My_chat_member struct {
			Chat TgResp_Chat `json:"chat"`
		} `json:"my_chat_member"`
	}
}

type TgResp_Chat struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Is_forum   bool   `json:"is_forum"`
}

type TgResp_User struct {
	ID                          int    `json:"id"`
	Is_bot                      bool   `json:"is_bot"`
	First_name                  string `json:"first_name"`
	Last_name                   string `json:"last_name"`
	Username                    string `json:"username"`
	Language_code               string `json:"language_code"`
	Is_premium                  bool   `json:"is_premium"`
	Added_to_attachment_menu    bool   `json:"added_to_attachment_menu"`
	Can_join_groups             bool   `json:"can_join_groups"`
	Can_read_all_group_messages bool   `json:"can_read_all_group_messages"`
	Supports_inline_queries     bool   `json:"supports_inline_queries"`
}

type TgResp_Message struct {
	Message_id        int         `json:"message_id"`
	Message_thread_id int         `json:"message_thread_id"`
	From              TgResp_User `json:"from"`
	Sender_chat       TgResp_Chat `json:"sender_chat"`
	Date              int         `json:"date"`
	Chat              TgResp_Chat `json:"chat"`
	Forward_from      TgResp_User `json:"forward_from"`
	Forward_from_chat TgResp_Chat `json:"forward_from_chat"`

	Text string `json:"text"`
}

type TgResp_Update struct {
	Update_id           int            `json:"update_id"`
	Message             TgResp_Message `json:"message"`
	Edited_message      TgResp_Message `json:"edited_message"`
	Channel_post        TgResp_Message `json:"channel_post"`
	Edited_channel_post TgResp_Message `json:"edited_channel_post"`
}
