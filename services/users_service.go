package services

import (
	"go_users_api/domain/users"
	"go_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
