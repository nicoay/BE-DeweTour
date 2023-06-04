package authdto

type AuthRequest struct {
	Name     string `json:"fullName" validate:"required"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Role     string `json:"role" validate:"required"`
}
