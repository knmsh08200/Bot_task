package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/knmsh08200/Bot_task/internal/handler"
)

func BotWork(bot *tgbotapi.BotAPI, update tgbotapi.CallbackQuery, TicketHandler handler.BotTicketHandlers) {

	if update.Message != nil { // Если это сообщение
		TicketHandler.TicketHandler(bot, update.Message)
	} else if update.CallbackQuery != nil { // Если это callback query
		TicketHandler.HandleCallback(bot, update.CallbackQuery)
	}

}
