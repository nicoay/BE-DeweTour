package handlers

import (
	authdto "dumbmerch/dto/auth"
	resultdto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/pkg/bcrypt"
	jwtToken "dumbmerch/pkg/jwt"
	"dumbmerch/repository"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Address:  request.Address,
		Role:     request.Role,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: "Email sama atau sudah terdaftar"})
	}
	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerate := jwtToken.GenerateToken(&claims)
	if errGenerate != nil {
		log.Println(errGenerate, "ini generae")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	registerResponse := authdto.RegisterResponse{
		Email:    user.Email,
		Token:    token,
		Response: data,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: registerResponse,
	})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ini bind error",
		})
	}
	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}
	// cek email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "ini salah email",
		})
	}
	// cek password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Wrong password",
		})
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerate := jwtToken.GenerateToken(&claims)
	if errGenerate != nil {
		log.Println(errGenerate, "ini generae")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: loginResponse,
	})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertAuthResponse(user)})
}

func convertAuthResponse(user models.User) authdto.CheckAuthResponse {
	return authdto.CheckAuthResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		Address:     user.Address,
		Gender:      user.Gender,
		Transaction: user.Transaction,
		Role:        user.Role,
	}
}
