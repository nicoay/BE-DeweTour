package models

import "time"

type Profile struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string    `json:"phone" gorm:"type : varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// for relation with user

type ProfileResponse struct {
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	UserId  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
