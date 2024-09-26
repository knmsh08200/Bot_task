package models

type TicketResponse struct {
	UserID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type TicketRequest struct {
	TicketID    string `json:"id"`
	UserID      int    `json:"user_id"`
	Departament string `json:"departament"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}
