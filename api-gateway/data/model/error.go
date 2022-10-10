package model

import (

)

type BaseError struct {
	Message string `json:"message"`
}

func NewError(message string) BaseError {
	return BaseError{message}
}

func (e BaseError) Error() string {
	return e.Message
}