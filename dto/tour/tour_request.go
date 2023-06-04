package tourdto

type CreateTour struct {
	Title          string `json:"title" form:"title" validate:"required"`
	CountryID      int    `json:"country_id" form:"country_id" validate:"required"`
	Accomodation   string `json:"accomodation" form:"accomodation" gorm:"type:varchar(255)" validate:"required"`
	Transportation string `json:"transport" form:"transport" gorm:"type:varchar(255)" validate:"required"`
	Eat            string `json:"eat" from:"eat" gorm:"type:varchar(255)" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" gorm:"type:varchar(255)" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	QuotaCurrent   int    `json:"quota_current" form:"quota_current" validate:"required"`
	Desc           string `json:"description" form:"description" gorm:"type:varchar(255)" validate:"required"`
	Image          string `json:"image" form:"image" gorm:"type:varchar(255)" validate:"required"`
}

type UpdateTour struct {
	Title          string `json:"title" form:"title" validate:"required"`
	CountryID      int    `json:"country_id" form:"country_id" validate:"required"`
	Accomodation   string `json:"accomodation" form:"accomodation" gorm:"type:varchar(255)" validate:"required"`
	Transportation string `json:"transport" form:"transport" gorm:"type:varchar(255)" validate:"required"`
	Eat            string `json:"eat" from:"eat" gorm:"type:varchar(255)" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" gorm:"type:varchar(255)" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	QuotaCurrent   int    `json:"quota_current" form:"quota_current" validate:"required"`
	Desc           string `json:"description" form:"description" gorm:"type:varchar(255)" validate:"required"`
	Image          string `json:"image" form:"image" gorm:"type:varchar(255)" validate:"required"`
}
