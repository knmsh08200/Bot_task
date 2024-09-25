package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotTicketHandlers interface {
	HandleCallback(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery)
	TicketHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message)
}
