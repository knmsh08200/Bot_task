package broker

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/knmsh08200/Bot_task/internal/models"
)

type ticketRedis struct { // change name
	rdb *redis.Client
}

func NewRepRedis(r *redis.Client) *ticketRedis {
	return &ticketRedis{rdb: r}
}

func (t *ticketRedis) CreateTicket(ctx context.Context, request models.TicketRequest) (models.TicketResponse, error) {

	ticketResp := models.TicketResponse{
		ID:     request.ID,
		Title:  request.Title,
		Body:   request.Body,
		Status: "новый",
	}

	// Кэшируем тикет в Redis
	err := t.rdb.Set(ctx, strconv.Itoa(request.ID), ticketResp, 0).Err() // set expiratiom time
	if err != nil {
		return ticketResp, err
	}

	// Добавляем тикет в список для уведомления администраторов
	adminListKey := request.Title + "_admin_notifications"
	err = t.rdb.RPush(ctx, adminListKey, request.ID).Err()
	if err != nil {
		return ticketResp, err
	}

	return ticketResp, nil

}
