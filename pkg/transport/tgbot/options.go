package tgbot

import (
	"context"
)

func (b *psycBot) run() {
	updates := b.bot.ListenForWebhook("/")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// b.bot.Send(menu)

	}
}

func (b *psycBot) Run(ctx context.Context) {
	go b.run()
	<-ctx.Done()
}
