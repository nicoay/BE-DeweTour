package tourdto

import "dumbmerch/models"

type TourResponse struct {
	Title          string                 `json:"title" `
	CountryID      int                    `json:"country_id" `
	Countries      models.CountryResponse `json:"country" `
	Accomodation   string                 `json:"accomodation" `
	Transportation string                 `json:"transport" `
	Eat            string                 `json:"eat" `
	Day            int                    `json:"day"`
	Night          int                    `json:"night" `
	DateTrip       string                 `json:"date_trip" `
	Price          int                    `json:"price" `
	Quota          int                    `json:"quota" `
	QuotaCurrent   int                    `json:"quota_current" `
	Desc           string                 `json:"description"`
	Image          string                 `json:"image"`
}
