package handlers

import (
	countrydto "dumbmerch/dto/country"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"

	// "time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCountry struct {
	CountryRepository repository.CountryRepository
}

func HandleCountry(CountryRepository repository.CountryRepository) *handlerCountry {
	return &handlerCountry{CountryRepository}
}

func (h *handlerCountry) FindCountries(c echo.Context) error {
	users, err := h.CountryRepository.FindCountries()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users,
	})
}
func (h *handlerCountry) GetCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users,
	})
}


func (h *handlerCountry) CreateCountry(c echo.Context) error {
	request := new(countrydto.CreateCountry)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	country := models.Country{
		Name: request.Name,
	}

	data, err := h.CountryRepository.CreateCountry(country)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertCountryResponse(data)})
}

func (h *handlerCountry) UpdateCountry(c echo.Context) error {
	request := new(countrydto.UpdateCountry)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.Name != "" {
		country.Name = request.Name
	}

	data, err := h.CountryRepository.UpdateCountry(country)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertCountryResponse(data)})
}

func (h *handlerCountry) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data, err := h.CountryRepository.DeleteCountry(id, country)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertCountryResponse(data),
	})

}

func convertCountryResponse(country models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		ID:   country.ID,
		Name: country.Name,
	}
}

// func GetPeople(c echo.Context) error {
// 	c.Response().Header().Set("Content-type", "application/json")
// 	id := c.Param("id")

// 	var Data People
// 	var cekId = false

// 	for _, talent := range Talent {
// 		if id == talent.Id {
// 			cekId = true
// 			Data = talent
// 		}
// 	}

// 	if !cekId {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID " + id + "not found")
// 	}

// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(Data)

// }

// func DeletePeople(c echo.Context) error {
// 	id := c.Param("id")
// 	var index = 0
// 	var cekId = false

// 	for i, talent := range Talent {
// 		if id == talent.Id {
// 			cekId = true
// 			index = i
// 		}
// 	}
// 	if !cekId {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID " + id + "not found")
// 	}
// 	Talent = append(Talent[:index], Talent[index+1:]...)
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(Talent)
// }
