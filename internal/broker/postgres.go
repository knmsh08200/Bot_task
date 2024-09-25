package broker

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/knmsh08200/Bot_task/internal/models"
)

type ticketPostgres struct {
	db *sql.DB
}

func NewRepPostgres(database *sql.DB) *ticketPostgres {
	return &ticketPostgres{db: database}
}

func (t *ticketPostgres) CreateTicket(ctx context.Context, request models.TicketRequest) (int, error) {
	id := 0
	err := t.db.QueryRowContext(ctx, "INSERT INTO tickets (id,department,title,body) VALUES ($1, $2, $3, $4) RETURNING id", request.ID, request.Departament, request.Title, request.Body).Scan(&id) // можно создать миграции для таблицы tickets
	if err != nil {
		log.Printf("Не удалось внести тикет в постгрес")
		return 0, err
	}
	return id, err
}

func (t *ticketPostgres) GetTicket(ctx context.Context, ticketID int) (models.TicketResponse, error) {
	query := `
	SELECT id,title,body FROM tickets WHERE id = $1
`
	//можно по разному искать тикет, здесь по id
	var ticket models.TicketResponse
	err := t.db.QueryRowContext(ctx, query, ticketID).Scan(&ticket.ID, &ticket.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return ticket, fmt.Errorf("article not found")
		}
		return ticket, err
	}

	return ticket, nil // Доделать логику вставки
}
