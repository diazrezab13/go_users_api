package success

import (
	"net/http"
)

type RestSuccess struct {
	Message string      `json:"message"`
	Status  int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewCreateSuccess(message string, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Status:  http.StatusCreated,
		Data:    data,
	}
}

func NewGetUsersSuccess(message string, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Status:  http.StatusOK,
		Data:    data,
	}
}
