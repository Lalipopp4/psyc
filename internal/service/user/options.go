package user

import (
	"context"
	"psyc/internal/errors"
	"psyc/internal/models"

	"psyc/pkg/scripts"
)

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	id, pass := s.repo.GetIDPassword(ctx, email)
	if !scripts.CheckPasswordHash(password, pass) {
		return "", errors.ErrorNotFound{Msg: errors.ErrUserLogin}
	}
	token, err := scripts.GenerateJWTUserToken(id, email)
	if err != nil {
		return "", err
	}
	if err := s.cache.Add(ctx, id, email); err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Register(ctx context.Context, user *models.User) (string, error) {
	if len(user.Password) < 8 {
		return "", errors.ErrorData{Msg: "password must be at least 8 characters"}
	}
	user.ID = scripts.GenerateID()
	token, err := scripts.GenerateJWTUserToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, s.repo.Add(ctx, user)
}
