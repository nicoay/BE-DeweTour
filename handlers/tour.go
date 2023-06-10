package handlers

import (
	dto "dumbmerch/dto/result"
	tourdto "dumbmerch/dto/tour"
	"dumbmerch/models"
	"dumbmerch/repository"
	"fmt"

	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTour struct {
	TourRepository repository.TourRepository
}

func HandleTour(TourRepository repository.TourRepository) *handlerTour {
	return &handlerTour{TourRepository}
}

var path_file = "http://localhost:5000/uploads/"

func (h *handlerTour) FindTours(c echo.Context) error {
	Tours, err := h.TourRepository.FindTours()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	for i, p := range Tours {
		Tours[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: Tours,
	})
}

func (h *handlerTour) GetTour(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tour, err := h.TourRepository.GetTour(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	tour.Image = path_file + tour.Image
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: tour,
	})
}
func (h *handlerTour) GetCountryTour(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tour, err := h.TourRepository.GetCountryTour(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: tour,
	})
}

func (h *handlerTour) CreateTour(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	countryId, _ := strconv.Atoi(c.FormValue("country_id"))
	day, _ := strconv.Atoi(c.FormValue("day"))
	night, _ := strconv.Atoi(c.FormValue("night"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	quota, _ := strconv.Atoi(c.FormValue("quota"))
	quotaCurrent, _ := strconv.Atoi(c.FormValue("quota_current"))

	request := tourdto.CreateTour{
		Title:          c.FormValue("title"),
		CountryID:      countryId,
		Accomodation:   c.FormValue("accomodation"),
		Transportation: c.FormValue("transport"),
		Eat:            c.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       c.FormValue("date_trip"),
		Price:          price,
		Quota:          quota,
		QuotaCurrent:   quotaCurrent,
		Desc:           c.FormValue("description"),
		Image:          dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	datas, err := h.TourRepository.GetCountryTour(request.CountryID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	fmt.Println(datas)

	tour := models.Tour{
		Title:          request.Title,
		CountryID:      request.CountryID,
		Countries:      datas,
		Accomodation:   request.Accomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		QuotaCurrent:   request.QuotaCurrent,
		Desc:           request.Desc,
		Image:          request.Image,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	data, err := h.TourRepository.CreateTour(tour)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Success Add Data",
		Data:    data,
	})
}
func (h *handlerTour) DeleteTour(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Tour, err := h.TourRepository.GetTour(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data, err := h.TourRepository.DeleteTour(id, Tour)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTourResponse(data),
	})

}
func (h *handlerTour) UpdateTour(c echo.Context) error {
	countryId, _ := strconv.Atoi(c.FormValue("country_id"))
	day, _ := strconv.Atoi(c.FormValue("day"))
	night, _ := strconv.Atoi(c.FormValue("night"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	quota, _ := strconv.Atoi(c.FormValue("quota"))
	quotaCurrent, _ := strconv.Atoi(c.FormValue("quota_current"))
	request := tourdto.UpdateTour{
		Title:          c.FormValue("title"),
		CountryID:      countryId,
		Accomodation:   c.FormValue("accomodation"),
		Transportation: c.FormValue("transport"),
		Eat:            c.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       c.FormValue("date_trip"),
		Price:          price,
		Quota:          quota,
		QuotaCurrent:   quotaCurrent,
		Desc:           c.FormValue("description"),
	}

	id, _ := strconv.Atoi(c.Param("id"))

	// datas, err := h.TourRepository.GetCountryTour(request.CountryID)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	})
	// }

	tour, err := h.TourRepository.GetTour(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	quotas := tour.Quota

	if quotaCurrent > quotas {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "QuotaCurrent cannot exceed Quota",
		})
	}

	if quotas <= tour.QuotaCurrent {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Quota already got Limit",
		})
	}
	if tour.QuotaCurrent+request.QuotaCurrent > quotas {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Quota not enough ,please decrease your order",
		})
	}
	if request.QuotaCurrent != 0 {
		tour.QuotaCurrent = tour.QuotaCurrent + request.QuotaCurrent
	}
	// fmt.Println(tour.QuotaCurrent)

	// if request.Title != "" {
	// 	tour.Title = request.Title
	// }

	// if request.CountryID != 0 {
	// 	tour.CountryID = request.CountryID
	// }

	// tour.Countries = datas

	// if request.Accomodation != "" {
	// 	tour.Accomodation = request.Accomodation
	// }

	// if request.Transportation != "" {
	// 	tour.Transportation = request.Transportation
	// }
	// if request.Eat != "" {
	// 	tour.Eat = request.Eat
	// }
	// if request.Day != 0 {
	// 	tour.Day = request.Day
	// }
	// if request.Night != 0 {
	// 	tour.Night = request.Night
	// }
	// if request.DateTrip != "" {
	// 	tour.DateTrip = request.DateTrip
	// }
	// if request.Price != 0 {
	// 	tour.Price = request.Price
	// }
	// if request.Quota != 0 {
	// 	tour.Quota = request.Quota
	// }
	// if request.Desc != "" {
	// 	tour.Desc = request.Desc
	// }
	// if request.Image != "" {
	// 	tour.Image = request.Image
	// }

	data, err := h.TourRepository.UpdateTour(tour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertTourResponse(data)})
}

func convertTourResponse(tour models.Tour) tourdto.TourResponse {
	return tourdto.TourResponse{
		Title:          tour.Title,
		CountryID:      tour.CountryID,
		Countries:      tour.Countries,
		Accomodation:   tour.Accomodation,
		Transportation: tour.Transportation,
		Eat:            tour.Eat,
		Day:            tour.Day,
		Night:          tour.Night,
		DateTrip:       tour.DateTrip,
		Price:          tour.Price,
		Quota:          tour.Quota,
		QuotaCurrent:   tour.QuotaCurrent,
		Desc:           tour.Desc,
		Image:          tour.Image,
	}
}
