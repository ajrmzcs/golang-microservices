package controllers

import (
	"github.com/ajrmzcs/golang-microservices/mvc/services"
	"github.com/ajrmzcs/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad request",
		}
		utils.RespondError(c, apiErr)
		//c.JSON(apiErr.Status, apiErr)
		//jsonValue, _ := json.Marshal(apiErr)
		//res.WriteHeader(apiErr.Status)
		//_, _ = res.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		//c.JSON(apiErr.Status, apiErr)
		//jsonValue, _ := json.Marshal(apiErr)
		//res.WriteHeader(apiErr.Status)
		//_, _ = res.Write([]byte(jsonValue))
		return
	}

	utils.Respond(c, http.StatusOK, user)

	//c.JSON(http.StatusOK, user)
	//jsonValue, _ := json.Marshal(user)
	//_, _ = res.Write([]byte(string(jsonValue)))
}