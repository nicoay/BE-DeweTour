package handlers

import (
	dto "dumbmerch/dto/result"
	transdto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repository"
	"time"

	// "fmt"

	// "fmt"
	"net/http"
	"strconv"

	// "time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repository.TransactionRepository
}

func HandleTransaction(TransactionRepository repository.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	Transactions, err := h.TransactionRepository.FindTransactions()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: Transactions,
	})
}
func (h *handlerTransaction) GetTourById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tour, err := h.TransactionRepository.GetTourById(id)

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

func (h *handlerTransaction) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.TransactionRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: user,
	})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: Transaction,
	})
}

// func (h *handlerTransaction) GetCountryTransaction(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	Transaction, err := h.TransactionRepository.GetCountryTransaction(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, dto.SuccessResult{
// 		Code: http.StatusOK,
// 		Data: Transaction,
// 	})
// }

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	countryQty, _ := strconv.Atoi(c.FormValue("counter_qty"))
	total, _ := strconv.Atoi(c.FormValue("total"))
	tourId, _ := strconv.Atoi(c.FormValue("tour_id"))

	request := transdto.CreateTransaction{
		CounterQty: countryQty,
		Total:      total,
		Status:     c.FormValue("status"),
		Attachment: dataFile,
		TourID:     tourId,
	}
	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// datas, err := h.TransactionRepository.GetCountryTransaction(request.CountryID)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	})
	// }
	// fmt.Println(datas)
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	tour, err := h.TransactionRepository.GetTourById(request.TourID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := h.TransactionRepository.GetUser(int(userId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	Transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TourID:     request.TourID,
		Tour:       tour,
		UserID:     int(userId),
		User:       user,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	data, err := h.TransactionRepository.CreateTransaction(Transaction)

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

func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	request := new(transdto.UpdateTransaction)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	// datas, err := h.TransactionRepository.GetCountryTransaction(request.CountryID)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	})
	// }

	Transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// Transaction.Countries = datas

	if request.CounterQty != 0 {
		Transaction.CounterQty = request.CounterQty
	}
	if request.Total != 0 {
		Transaction.Total = request.Total
	}
	if request.Status != "" {
		Transaction.Status = request.Status
	}
	if request.Attachment != "" {
		Transaction.Attachment = request.Attachment
	}
	if request.TourID != 0 {
		Transaction.TourID = request.TourID
	}
	if request.UserID != 0 {
		Transaction.UserID = request.UserID
	}

	// datas, err := h.TransactionRepository.GetCountryTransaction(id)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	data, err := h.TransactionRepository.UpdateTransaction(Transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertTransactionResponse(data)})
}
func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data, err := h.TransactionRepository.DeleteTransaction(id, Transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTransactionResponse(data),
	})

}

func convertTransactionResponse(Transaction models.Transaction) transdto.TransactionResponse {
	return transdto.TransactionResponse{
		CounterQty: Transaction.CounterQty,
		Total:      Transaction.Total,
		Status:     Transaction.Status,
		Attachment: Transaction.Attachment,
		TourID:     Transaction.TourID,
		UserID:     Transaction.UserID,
	}
}
