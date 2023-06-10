package models

import "time"

type Transaction struct {
	ID         int                  `json:"id" form:"id" gorm:"primary_key"`
	CounterQty int                  `json:"counter_qty" form:"counter_qty"`
	Total      int                  `json:"total" form:"total"`
	Status     string               `json:"status" form:"status"`
	UserID     int                  `json:"id_user"`
	User       UsersProfileResponse `json:"user" form:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TourID     int                  `json:"id_tour"`
	Tour       TourResponse         `json:"tour" form:"user" gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"-"`
}
type TransactionResponse struct {
	ID         int          `json:"id"`
	CounterQty int          `json:"counter_qty"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	UserID     int          `json:"-"`
	TourID     int          `json:"id_tour"`
	Tour       TourResponse `json:"tour"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"-"`
}
func (TransactionResponse) TableName() string {
	return "transactions"
}
