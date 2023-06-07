package countrydto

type CreateCountry struct {
	Name string `json:"name" form:"name" validate:"required"`
}
type UpdateCountry struct {
	Name string `json:"name" form:"name" validate:"required"`
}
