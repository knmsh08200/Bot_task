package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotTicketHandlers interface {
	HandleCallback(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery)
	TicketHandlerTicketHandler(ctx context.Context, bot *tgbotapi.BotAPI, message *tgbotapi.Message)
}
