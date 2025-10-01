package services

import (
	"fmt"
	"net/smtp"
	"finance-app/backend/internal/entity"
)

type EmailService struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	fromEmail    string
}

func NewEmailService(host, port, username, password, from string) *EmailService {
	return &EmailService{
		smtpHost:     host,
		smtpPort:     port,
		smtpUsername: username,
		smtpPassword: password,
		fromEmail:    from,
	}
}

func (e *EmailService) SendReminder(bill *entity.Bill, user *entity.User) error {
	if !user.EmailNotify {
		return fmt.Errorf("email notification disabled")
	}

	subject := fmt.Sprintf("Bill Reminder: %s", bill.Name)
	body := fmt.Sprintf(
		"Hi %s,\n\n"+
			"This is a reminder for your upcoming bill:\n\n"+
			"Bill Name: %s\n"+
			"Amount: Rp%.2f\n"+
			"Due Date: %s\n\n"+
			"Please make sure to pay on time to avoid late fees.\n\n"+
			"Best regards,\n"+
			"Finance App Team",
		user.Name,
		bill.Name,
		bill.Amount,
		bill.DueDate.Format("2006-01-02"),
	)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		e.fromEmail, user.Email, subject, body)

	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpHost)
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)

	err := smtp.SendMail(addr, auth, e.fromEmail, []string{user.Email}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}