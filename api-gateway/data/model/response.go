package model

type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type SuccessResponse struct {
	Timestamp string `json:"timestamp"`
	Message string `json:"message"`
}