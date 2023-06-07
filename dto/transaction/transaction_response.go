package transdto

type TransactionResponse struct {
	CounterQty int    `json:"counter_qty"`
	Total      int    `json:"total" `
	Status     string `json:"status" `
	Attachment string `json:"attachment" `
	TourID     int    `json:"tour" `
	UserID     int    `json:"user"`
}
