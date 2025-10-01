package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"finance-app/backend/internal/entity"
)

type TelegramService struct {
	botToken string
}

func NewTelegramService(botToken string) *TelegramService {
	return &TelegramService{botToken: botToken}
}

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func (t *TelegramService) SendReminder(bill *entity.Bill, user *entity.User) error {
	if user.TelegramChatID == "" || !user.TelegramNotify {
		return fmt.Errorf("telegram notification disabled or chat ID not set")
	}

	message := fmt.Sprintf(
		"💰 *Bill Reminder*\n\n"+
			"*%s*\n"+
			"Amount: *Rp%.2f*\n"+
			"Due Date: *%s*\n"+
			"Status: *%s*\n\n"+
			"Don't forget to pay your bill! 💳",
		bill.Name,
		bill.Amount,
		bill.DueDate.Format("2006-01-02"),
		bill.Status,
	)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)

	msg := TelegramMessage{
		ChatID: user.TelegramChatID,
		Text:   message,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", string(body))
	}

	return nil
}

func (t *TelegramService) SendTestMessage(chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)

	msg := TelegramMessage{
		ChatID: chatID,
		Text:   message,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", string(body))
	}

	return nil
}