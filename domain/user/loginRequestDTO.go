package user

//this is our model for login request
type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}