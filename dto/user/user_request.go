package userdto

type CreateUser struct {
	Name     string `json:"fullName" form:"fullName" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"  binding:"required, email" gorm:"unique; not null"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"fullName" form:"fullName" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	Picture  string `json:"picture" form:"picture"`
}
