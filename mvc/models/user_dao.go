package models

import (
	"fmt"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[uint64]*User{
		123:&User{Id:1, FirstName:"Antonio", LastName:"Ramirez", Email:"ajrmzcs@gmail.com"},
		}
	UserDao userDaoInterface
)

// Need to clarify this
func init(){
	UserDao = &userDao{}
}

type userDaoInterface interface{
	GetUser(uint64)(*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not found",
	}
}