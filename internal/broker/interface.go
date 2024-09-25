package broker

import (
	"context"

	"github.com/knmsh08200/Bot_task/internal/models"
)

type TicketService interface {
	CreateTicket(ctx context.Context, request models.TicketRequest) (models.TicketResponse, error)
}

// type TicketServiceDB interface {
// 	CreateTicket(ctx context.Context, request models.TicketRequest) (int, error)
// 	GetTicket(ctx context.Context, ticketID int) (models.TicketResponse, error)
// }
