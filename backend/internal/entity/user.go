package entity

import (
	"time"
)

type User struct {
	ID             uint      `json:"id"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	Name           string    `json:"name"`
	TelegramChatID string    `json:"telegram_chat_id"`
	EmailNotify    bool      `json:"email_notify"`
	TelegramNotify bool      `json:"telegram_notify"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
