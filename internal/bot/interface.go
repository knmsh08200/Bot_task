package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotUpdates interface {
	GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type BotRepository interface {
	NotifyAdmins(ctx context.Context, bot *tgbotapi.BotAPI)
}
