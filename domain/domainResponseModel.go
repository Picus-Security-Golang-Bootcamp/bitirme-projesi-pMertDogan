package domain

type ResponseModel struct {
	ErrMsg       string `json:"errMsg,omitempty"`
	ErrDsc       string `json:"errDsc,omitempty"`
	ResponseCode int    `json:"responseCode"`

	Data interface{} `json:"data,omitempty"`
}
