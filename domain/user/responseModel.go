package user

type ResponseModel struct{
	ErrMsg string `json:"errMsg,omitempty"` //if errMsg is empty, it means there is no error. With the help of the omitempty, we can omit the field from the response
	ResponseCode int
	AccesToken string `json:"accessToken,omitempty"`
	RefreshToken string	`json:"refreshToken,omitempty"`
}