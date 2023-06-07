package models

import "time"

type Tour struct {
	ID             int             `json:"id_tour" gorm:"primary_key:auto_increment"`
	Title          string          `json:"title" form:"title" gorm:"type:varchar(255) "`
	CountryID      int             `json:"id_country" form:"id_country"`
	Countries      CountryResponse `json:"country" form:"country" gorm:"foreignKey:CountryID;constraint:OnUpdate:SET_DEFAULT,OnDelete:CASCADE;"`
	Accomodation   string          `json:"accomodation" form:"accomodation" gorm:"type:varchar(255)"`
	Transportation string          `json:"transport" form:"transport" gorm:"type:varchar(255)"`
	Eat            string          `json:"eat" from:"eat" gorm:"type:varchar(255)"`
	Day            int             `json:"day" form:"day"`
	Night          int             `json:"night" form:"night"`
	DateTrip       string          `json:"date_trip" form:"date_trip" gorm:"type:varchar(255)"`
	Price          int             `json:"price" form:"price"`
	Quota          int             `json:"quota" form:"quota"`
	QuotaCurrent   int             `json:"quota_current" form:"quota_current"`
	Desc           string          `json:"description" form:"description" gorm:"type:varchar(255)"`
	Image          string          `json:"image" form:"image" gorm:"type:varchar(255)"`
	CreatedAt      time.Time       `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
}

type TourResponse struct {
	ID             int             `json:"id_tour"`
	Title          string          `json:"title"`
	CountryID      int             `json:"id_country" form:"id_country"`
	Countries      CountryResponse `json:"country" form:"country" gorm:"foreignKey:CountryID"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transport"`
	Eat            string          `json:"eat"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       string          `json:"date_trip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota" `
	Desc           string          `json:"description"`
	Image          string          `json:"image"`
}

func (TourResponse) TableName() string {
	return "tours"
}
