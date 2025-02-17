package service

import (
	"fmt"
	"net/smtp"
)

// EmailService defines the interface for sending emails.
type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
	from     string
}

// NewEmailService creates a new instance of emailService with the provided SMTP settings.
func NewEmailService(smtpHost, smtpPort, username, password, from string) EmailService {
	return &emailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
		from:     from,
	}
}

// SendEmail sends a real email using net/smtp.
func (s *emailService) SendEmail(to, subject, body string) error {
	// Construct the email message.
	msg := []byte("From: " + s.from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Set up authentication information.
	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	// Send the email.
	if err := smtp.SendMail(addr, auth, s.from, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
