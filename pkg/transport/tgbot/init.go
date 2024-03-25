package tgbot

import (
	"psyc/internal/mock/service/result"
	"psyc/internal/service/user"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type psycBot struct {
	result result.Service
	user   user.Service

	bot *tgbotapi.BotAPI
}
