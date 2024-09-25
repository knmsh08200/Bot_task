package models

type TicketResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status string `json:"status"` // под вопросом данное поле
}

type TicketRequest struct {
	ID          int    `json:"id"`
	Departament string `json:"departament"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}
