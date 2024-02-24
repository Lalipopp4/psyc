package result

import (
	"psyc/internal/controllers/db/result"
	"psyc/internal/controllers/mail"
)

type resultService struct {
}

func New(repo result.Repository, mail mail.MailSender) Service {
	return &resultService{}
}
