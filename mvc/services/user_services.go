package services

import (
	"github.com/ajrmzcs/golang-microservices/mvc/models"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
)

func GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userId)
}