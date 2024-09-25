package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAPI struct {
	Bot *tgbotapi.BotAPI
}

func (tb *TelegramAPI) GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	return tb.Bot.GetUpdatesChan(u)
}

func NewTelegramBot(token string, debug bool) (*TelegramAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug
	return &TelegramAPI{Bot: bot}, nil
}

func StartBot(bot BotUpdates, timeout int) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates := bot.GetUpdatesChan(u)
	return updates, nil
}
