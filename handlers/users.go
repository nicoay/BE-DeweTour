package handlers

import (
	dto "dumbmerch/dto/result"
	userdto "dumbmerch/dto/user"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repository.UserRepository
}

func HandleUser(UserRepository repository.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

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

func (h *handlerUser) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserById(id)

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

func (h *handlerUser) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUser)
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

	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.UserRepository.CreateUser(user)

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
func (h *handlerUser) UpdateUser(c echo.Context) error {
	request := new(userdto.UpdateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Address != "" {
		user.Address = request.Address
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handlerUser) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data, err := h.UserRepository.DeleteUser(id, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(data),
	})

}

func convertResponse(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
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
