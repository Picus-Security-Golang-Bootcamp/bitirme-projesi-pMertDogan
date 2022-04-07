package category

type ResponseModel struct{
	ErrMsg string `json:"errMsg,omitempty"`
	ResponseCode int `json:"responseCode"`
	Data Categorys `json:"data,omitempty"`
}