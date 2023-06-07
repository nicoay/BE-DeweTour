package userdto

// import "dumbmerch/models"

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Country  models.Country `json:"country"`
}
