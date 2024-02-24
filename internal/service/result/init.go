package result

import (
	"psyc/internal/controllers/db/result"
	"psyc/internal/controllers/mail"
)

type resultService struct {
	repo result.Repository
	mail mail.MailSender
}

func New(repo result.Repository, mail mail.MailSender) Service {
	return &resultService{repo, mail}
}
