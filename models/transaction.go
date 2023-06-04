package models

import "time"

type Transaction struct {
	ID         int                  `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	CounterQty int                  `json:"counter_qty" form:"counter_qty"`
	Total      int                  `json:"total" form:"total"`
	Status     string               `json:"status" form:"status"`
	Attachment string               `json:"attachment" form:"attachment"`
	UserID     int                  `json:"id_user"`
	User       UsersProfileResponse `json:"user" form:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TourID     int                  `json:"id_tour"`
	Tour       TourResponse         `json:"tour" form:"user" gorm:"foreignKey:TourID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time            `json:"-"`
	UpdatedAt  time.Time            `json:"-"`
}
type TransactionResponse struct {
	ID         int          `json:"id"`
	CounterQty int          `json:"counter_qty"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	Attachment string       `json:"attachment"`
	UserID     int          `json:"-"`
	TourID     int          `json:"-"`
	Tour       TourResponse `json:"tour"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
