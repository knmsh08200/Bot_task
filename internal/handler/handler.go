package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/knmsh08200/Bot_task/internal/broker"
	"github.com/knmsh08200/Bot_task/internal/models"
)

type TicketHandler struct {
	Service broker.TicketService
}

func (t *TicketHandler) TicketHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.Chat.ID

	switch message.Text {
	case "/start":
		t.handleStart(bot, userID)
	case "/new":
		t.handleNewTicket(bot, userID)
	case "IT", "Billing", "Support":
		t.handleSetTicket(bot, userID, message.Text)
	default:
		t.handleSetTicket(bot, userID, message.Text)
	}
}

func (t *TicketHandler) handleStart(bot *tgbotapi.BotAPI, userID int64) {
	msg := tgbotapi.NewMessage(userID, "Выберите подразделение: IT, Billing, Support.")
	bot.Send(msg)
}

func (t *TicketHandler) handleNewTicket(bot *tgbotapi.BotAPI, userID int64) {
	msg := tgbotapi.NewMessage(userID, "Выберите  подразделение: 1) Support 2)IT 3) Billing")
	bot.Send(msg)
}

func (t *TicketHandler) handleSetTicket(bot *tgbotapi.BotAPI, userID int64, input string) {

}

func handleFastAnswer(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world")
	bot.Send(msg)

	userId := update.Message.From.ID // id пользователя, который прислал сообщение

	// Создание кнопки быстрого ответа

	// Обработка нажатий на кнопку
	if update.CallbackQuery != nil {
		callback := update.CallbackQuery
		if callback.Data == "reply:"+string(userId) {
			// Отправка сообщения "Hello world" клиенту
			replyMsg := tgbotapi.NewMessage(callback.From.ID, "Hello world")
			bot.Send(replyMsg)

			// Ответ на callback
			callbackResponse := tgbotapi.NewCallback(callback.ID, "Message sent")
			bot.Send(callbackResponse)
		}
	}
}

func (t *TicketHandler) CreateTicket(ctx context.Context, request models.TicketRequest) (models.TicketResponse, error) {
	return t.Service.CreateTicket(ctx, request)
}
