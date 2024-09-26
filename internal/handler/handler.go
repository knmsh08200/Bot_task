package handler

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/knmsh08200/Bot_task/internal/bot"
	"github.com/knmsh08200/Bot_task/internal/models"
)

type TicketHandler struct {
	Provider bot.BotService
}

func NewTicketHandler(s bot.BotService) *TicketHandler {
	return &TicketHandler{Provider: s}
}

func (t *TicketHandler) TicketHandler(ctx context.Context, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.Chat.ID
	tickRes := t.cacheTicket()

	switch message.Text {
	case "/start":
		t.handleStart(bot, userID)
	case "/new":
		t.handleNewTicket(bot, userID)
	case "IT", "Billing", "Support":
		t.handleSetTicket(ctx, bot, userID, message.Text, tickRes)
	default:
		t.handleSetTicket(ctx, bot, userID, message.Text, tickRes)
	}
}

func (t *TicketHandler) handleStart(bot *tgbotapi.BotAPI, userID int64) {
	msg := tgbotapi.NewMessage(userID, "Добрый день")
	bot.Send(msg)
}

func (t *TicketHandler) handleNewTicket(bot *tgbotapi.BotAPI, userID int64) {
	msg := tgbotapi.NewMessage(userID, "Выберите  подразделение: 1) Support 2)IT 3) Billing")
	bot.Send(msg)
}

func (t *TicketHandler) handleSetTicket(ctx context.Context, bot *tgbotapi.BotAPI, userID int64, input string, ticket *models.TicketRequest) {

	// В зависимости от этапа диалога, сохраняем разные данные.
	switch input {
	case "IT", "Billing", "Support":
		ticket.Departament = input // Сохраняем выбранный департамент
		msg := tgbotapi.NewMessage(userID, "Введите описание проблемы.")
		bot.Send(msg)
	default:
		if ticket.Title == "" {
			ticket.Title = input
		} else if ticket.Body == "" {
			ticket.Body = input
		}

		ticket.TicketID = generateUniqueID()
		ticket.UserID = int(userID)
		key, err := t.CreateTicket(ctx, ticket)

		if err != nil {
			msg := tgbotapi.NewMessage(userID, "Произошла ошибка при создании тикета.")
			bot.Send(msg)
			return
		}

		// Формируем текст ответа с полным содержанием тикета
		ticketText := fmt.Sprintf(
			"Тикет успешно создан!\n\nПодразделение: %s\nОписание: %s\nНомер тикета: %s",
			ticket.Departament,
			ticket.Title,
			ticket.Body,
		)

		// Уведомляем пользователя полным содержанием тикета
		msg := tgbotapi.NewMessage(userID, ticketText)
		bot.Send(msg)

		// уведомление админа
		t.Provider.NotifyAdmins(ctx, key, bot)

	}
}

func (t *TicketHandler) CreateTicket(ctx context.Context, request *models.TicketRequest) (string, error) {
	return t.Provider.CreateTicket(ctx, request)
}

func (t *TicketHandler) HandleCallback(bot *tgbotapi.BotAPI, update *tgbotapi.CallbackQuery) {
	if update.Data == "hello_world" {
		// Отправляем сообщение обратно пользователю
		msg := tgbotapi.NewMessage(update.From.ID, "Hello world")
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}

		// Удаляем уведомление о нажатой кнопке
		callbackMsg := tgbotapi.NewCallback(update.ID, "Кнопка нажата")
		_, err = bot.Send(callbackMsg)
		if err != nil {
			log.Printf("Ошибка при отправке callback: %v", err)
		}
	}

}

func (t *TicketHandler) fastAnswer(ctx context.Context, bot *tgbotapi.BotAPI, userID int64) {
	// Создаем сообщение с кнопкой
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отправить 'Hello world'", "hello_world"),
		),
	)
	msg := tgbotapi.NewMessage(userID, "Hello world")
	msg.ReplyMarkup = replyMarkup

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

func (t *TicketHandler) cacheTicket() *models.TicketRequest {
	return &models.TicketRequest{}
}

func generateUniqueID() string {
	return uuid.NewString()
}
