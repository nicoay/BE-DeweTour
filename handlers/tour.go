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
	"github.com/labstack/echo/v4"
)

type handlerTour struct {
	TourRepository repository.TourRepository
}

func HandleTour(TourRepository repository.TourRepository) *handlerTour {
	return &handlerTour{TourRepository}
}

func (h *handlerTour) FindTours(c echo.Context) error {
	Tours, err := h.TourRepository.FindTours()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
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
	request := new(tourdto.CreateTour)
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
func (h *handlerTour) UpdateTour(c echo.Context) error {
	request := new(tourdto.UpdateTour)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	datas, err := h.TourRepository.GetCountryTour(request.CountryID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	tour, err := h.TourRepository.GetTour(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		tour.Title = request.Title
	}

	if request.CountryID != 0 {
		tour.CountryID = request.CountryID
	}
	
	tour.Countries = datas

	if request.Accomodation != "" {
		tour.Accomodation = request.Accomodation
	}

	if request.Transportation != "" {
		tour.Transportation = request.Transportation
	}
	if request.Eat != "" {
		tour.Eat = request.Eat
	}
	if request.Day != 0 {
		tour.Day = request.Day
	}
	if request.Night != 0 {
		tour.Night = request.Night
	}
	if request.DateTrip != "" {
		tour.DateTrip = request.DateTrip
	}
	if request.Price != 0 {
		tour.Price = request.Price
	}
	if request.Quota != 0 {
		tour.Quota = request.Quota
	}
	if request.Desc != "" {
		tour.Desc = request.Desc
	}
	if request.Image != "" {
		tour.Image = request.Image
	}
	// datas, err := h.TourRepository.GetCountryTour(id)
	// fmt.Println(datas)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
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
		Desc:           tour.Desc,
		Image:          tour.Image,
	}
}
