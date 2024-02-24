package mail

import "net/smtp"

type mailManager struct {
	auth smtp.Auth
	addr string
	mail string
}

func New(cfg *Config) *mailManager {
	return &mailManager{
		smtp.PlainAuth("", cfg.Mail.mail, cfg.Mail.password, cfg.Mail.host),
		cfg.Mail.host + ":" + cfg.Mail.port,
		cfg.Mail.mail,
	}
}
