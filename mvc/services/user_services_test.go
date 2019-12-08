package services

import (
	"github.com/ajrmzcs/golang-microservices/mvc/models"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId uint64)(*models.User, *utils.ApplicationError)
)

func init(){
	models.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUserService_GetUser(t *testing.T) {
	getUserFunction = func(userId uint64) (*models.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message: "user 0 not found",
			Status:  http.StatusNotFound,
			Code:    "not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestUserService_GetUserNoError(t *testing.T) {
	getUserFunction = func(userId uint64) (*models.User, *utils.ApplicationError) {
		return &models.User{
			Id:        1,
		}, nil
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
}
