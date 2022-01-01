package controllers

import (
	"go_users_api/domain/users"
	"go_users_api/services"
	"go_users_api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
