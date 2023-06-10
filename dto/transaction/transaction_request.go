package transdto

type CreateTransaction struct {
	CounterQty int    `json:"counter_qty" form:"counter_qty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" gorm:"type:varchar(255)"`
	TourID     int    `json:"tour_id" form:"tour_id" validate:"required"`
	UserID     int    `json:"user_id" form:"user_id"`
}

type UpdateTransaction struct {
	CounterQty int    `json:"counter_qty" form:"counter_qty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" gorm:"type:varchar(255)" validate:"required"`
	TourID     int    `json:"tour_id" form:"tour_id" validate:"required"`
	UserID     int    `json:"user_id" form:"user_id"`
}
