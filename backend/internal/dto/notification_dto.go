package dto

type NotificationSettingsRequest struct {
	TelegramNotify  bool   `json:"telegram_notify"`
	EmailNotify     bool   `json:"email_notify"`
	TelegramChatID  string `json:"telegram_chat_id"`
}

type TestTelegramRequest struct {
	Message string `json:"message" binding:"required"`
}