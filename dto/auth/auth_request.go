package authdto

type AuthRequest struct {
	Name     string `json:"fullName" validate:"required"`
	Email    string `json:"email" form:"email" binding:"required, email" gorm:"unique; not null"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
