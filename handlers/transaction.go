package handlers

import (
	dto "dumbmerch/dto/result"
	transdto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repository"
	"log"
	"os"
	"time"

	// "fmt"

	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
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
	request := new(transdto.CreateTransaction)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
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

	request.Status = "pending"
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
		ID:         transactionId,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TourID:     request.TourID,
		Tour:       tour,
		UserID:     int(userId),
		User:       user,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	data, err := h.TransactionRepository.CreateTransaction(Transaction)
	fmt.Println(data.TourID, "ini data create")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.ID),
			GrossAmt: int64(data.Total),
		},

		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.User.Name,
			Email: data.User.Email,
			Phone: data.User.Phone,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: snapResp,
	})
}
func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudstatus := notificationPayload["fraud_status"].(string)
	transId := notificationPayload["order_id"].(string)

	trans_id, _ := strconv.Atoi(transId)
	fmt.Println("ini notif", notificationPayload)
	fmt.Println("ini notif", trans_id)
	transaction, _ := h.TransactionRepository.GetTransaction(trans_id)

	if transactionStatus == "capture" {
		if fraudstatus == "challange" {
			h.TransactionRepository.UpdateTransaction("pending", trans_id)
		} else if fraudstatus == "accept" {
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", trans_id)
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", trans_id)
	} else if transactionStatus == "deny" {
		h.TransactionRepository.UpdateTransaction("failed", trans_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.TransactionRepository.UpdateTransaction("failed", trans_id)
	} else if transactionStatus == "pending" {
		h.TransactionRepository.UpdateTransaction("pending", trans_id)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: notificationPayload,
	})
}

func SendMail(status string, transaction models.Transaction) {
	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "Dewe Tour <skadskuds.f@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var tourName = transaction.Tour.Title
		var total = strconv.Itoa(transaction.Total)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Success")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		  <head>
		  <meta charset="UTF-8" />
		  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
		  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
		  <title>Document</title>
		  <style>
			h1 {
			color: brown;
			}
		  </style>
		  </head>
		  <body>
		  <h2>Product payment :</h2>
		  <ul style="list-style-type:none;">
			<li>Name : %s</li>
			<li>Total payment: Rp.%s</li>
			<li>Status : <b>%s</b></li>
		  </ul>
		  </body>
		</html>`, tourName, total, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
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

		TourID: Transaction.TourID,
		UserID: Transaction.UserID,
	}
}
