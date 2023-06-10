package models

import "time"

type Country struct {
	ID        int       `json:"id_country" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name_country" binding:"required, name_country" gorm:"unique; not null;varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
