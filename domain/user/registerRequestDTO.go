package user

//this is our model for login request
type RegisterRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}