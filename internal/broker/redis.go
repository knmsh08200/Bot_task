package broker

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/knmsh08200/Bot_task/internal/models"
)

type ticketRedis struct { // change name
	rdb *redis.Client
}

func NewRepRedis(r *redis.Client) *ticketRedis {
	return &ticketRedis{rdb: r}
}

func (t *ticketRedis) CreateTicket(ctx context.Context, request *models.TicketRequest) (string, error) {

	ticketJSON, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	// Кэшируем тикет в Redis, используя TicketID в качестве ключа
	err = t.rdb.Set(ctx, request.TicketID, ticketJSON, 0).Err()
	if err != nil {
		return "", err
	}

	return request.TicketID, nil

}

func (t *ticketRedis) GetTicket(ctx context.Context, ticketID string) (models.TicketResponse, error) {
	var ticketResp models.TicketResponse

	// Извлекаем тикет из Redis
	ticketJSON, err := t.rdb.Get(ctx, ticketID).Result()
	if err != nil {
		return ticketResp, err
	}

	// Декодируем JSON в структуру
	err = json.Unmarshal([]byte(ticketJSON), &ticketResp)
	if err != nil {
		return ticketResp, err
	}

	return ticketResp, nil

}
