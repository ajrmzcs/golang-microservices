package services

import (
	"github.com/ajrmzcs/golang-microservices/mvc/models"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return models.UserDao.GetUser(userId)
}