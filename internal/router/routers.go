package router

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/knmsh08200/Bot_task/internal/handler"
)

func BotWork(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, ticketHandler *handler.TicketHandler) {

	if update.Message != nil { // Если это сообщение
		ticketHandler.TicketHandler(ctx, bot, update.Message)
	} else if update.CallbackQuery != nil { // Если это callback query
		ticketHandler.HandleCallback(bot, update.CallbackQuery)
	}

}
