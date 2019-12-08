package services

import (
	"github.com/ajrmzcs/golang-microservices/mvc/models"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
	"net/http"
)

type itemService struct {}

var ItemService itemService

func (i *itemService) getItems(itemId string)(*models.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "Implement me",
		Status:  http.StatusInternalServerError,
		Code:    "Method not implemented yet",
	}
}
