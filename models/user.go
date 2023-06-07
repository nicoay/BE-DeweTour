package models

import "time"

type User struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	Name        string                `json:"fullName" form:"fullName" gorm:"varchar(255)"`
	Email       string                `json:"email" form:"email" binding:"required, email" gorm:"unique; not null"`
	Password    string                `json:"password" form:"password" gorm:"varchar(255)"`
	Gender      string                `json:"gender" form:"gender" gorm:"varchar(255)"`
	Phone       string                `json:"phone" form:"phone" gorm:"varchar(255)"`
	Address     string                `json:"address" form:"address" gorm:"varchar(255)"`
	Role        string                `json:"role" form:"role" gorm:"varchar(255)"`
	Transaction []TransactionResponse `json:"transaction"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"fullName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
