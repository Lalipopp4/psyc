package result

import (
	"context"
	"psyc/internal/controllers/db/result"
	"psyc/internal/controllers/mail"
	"psyc/internal/models"
)

type resultService struct {
	repo result.Repository
	mail mail.MailSender

	tests map[string]func(context.Context, *models.Test) (string, error)
}

func New(repo result.Repository, mail mail.MailSender) Service {
	s := &resultService{
		repo: repo,
		mail: mail,
	}
	s.tests = map[string]func(context.Context, *models.Test) (string, error){
		"keirsey": s.keirsey,
		"hall":    s.hall,
		"bass":    s.bass,
		"eysenck": s.eysenck,
	}

	return s
}
