package transdto

import "time"

type TransactionResponse struct {
	CounterQty int       `json:"counter_qty"`
	Total      int       `json:"total" `
	Status     string    `json:"status" `
	TourID     int       `json:"tour" `
	UserID     int       `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
}
