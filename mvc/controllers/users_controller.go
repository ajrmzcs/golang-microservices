package controllers

import (
	"encoding/json"
	"github.com/ajrmzcs/golang-microservices/mvc/services"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseUint(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		_, _ = res.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		_, _ = res.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	_, _ = res.Write([]byte(string(jsonValue)))
}