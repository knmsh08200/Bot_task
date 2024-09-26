package bot

import (
	"context"
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/knmsh08200/Bot_task/internal/admin"
	"github.com/knmsh08200/Bot_task/internal/broker"
	"github.com/knmsh08200/Bot_task/internal/models"
)

type BotService struct {
	Service broker.TicketService
}

func NewRep(b broker.TicketService) BotService {
	return BotService{Service: b}
}

func (b *BotService) NotifyAdmins(ctx context.Context, ticketID string, bot *tgbotapi.BotAPI) {
	for {
		// Извлекаем тикеты для каждого отдела
		departments := []string{"Support", "IT", "Billing"}
		for _, department := range departments {

			ticketResponce, err := b.Service.GetTicket(ctx, ticketID)
			if err != nil {
				log.Printf("Bad with redis")
			}
			adminID, exists := admin.AdminIDs[department]
			if !exists {
				log.Printf("Не найден админ для департамента: %s", department)
				continue
			}

			ticketMessage, err := json.Marshal(ticketResponce)
			if err != nil {
				log.Printf("Ошибка при сериализации тикета: %v", err)
				continue
			}

			// Отправляем уведомление админу
			message := tgbotapi.NewMessage(int64(adminID), "Новый тикет:\nID:"+string(ticketMessage))
			_, err = bot.Send(message)
			if err != nil {
				log.Printf("Уведомление не отправлено")
			} else {
				log.Printf("Уведомление отправлено администратору %d о новом тикете: %d", adminID, ticketResponce.UserID)
			}

		}
	}
}

func (t *BotService) CreateTicket(ctx context.Context, request *models.TicketRequest) (string, error) { // тут можно интерфейс использовать или попробовать разнести интерфейсы постгрес и редис
	return t.Service.CreateTicket(ctx, request)
}
