package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotUpdates interface {
	GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}
