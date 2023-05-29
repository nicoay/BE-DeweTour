package models

import "time"

type User struct {
	ID          int                   `json:"id"`
	Name        string                `json:"fullName" form:"fullName" gorm:"varchar(255)"`
	Email       string                `json:"email" form:"email" gorm:"varchar(255)"`
	Password    string                `json:"password" form:"password" gorm:"varchar(255)"`
	Phone       string                `json:"phone" form:"phone" gorm:"varchar(255)"`
	Address     string                `json:"address" form:"address" gorm:"varchar(255)"`
	Transaction []TransactionResponse `json:"transaction"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"namefullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
