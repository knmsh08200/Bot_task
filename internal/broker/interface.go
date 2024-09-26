package broker

import (
	"context"

	"github.com/knmsh08200/Bot_task/internal/models"
)

type TicketService interface {
	GetTicket(ctx context.Context, ticketID string) (models.TicketResponse, error)
	CreateTicket(ctx context.Context, request *models.TicketRequest) (string, error) // тут можно интерфейс использовать или попробовать разнести интерфейсы постгрес и редис
}
