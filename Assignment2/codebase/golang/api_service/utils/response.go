package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"octcarp/sustech/cs328/a2/api/grpc/logclient"
	"octcarp/sustech/cs328/a2/api/models"
)

func ResponseLog(c *gin.Context, code int, message string) {
	logMsg := fmt.Sprintf("[%s] %s %s, response code: %d, message: %s",
		c.Request.Method,
		c.Request.URL.Path,
		c.ClientIP(),
		code,
		message,
	)
	if code == http.StatusInternalServerError {
		logclient.Warning(logMsg, "default_id")
	} else {
		logclient.Info(logMsg, "default_id")
	}
}

func SendCodeWithMessage(c *gin.Context, code int, message string) {
	ResponseLog(c, code, message)
	c.JSON(code, models.Message{
		Message: message,
	})
}

func SendNotFoundErr(c *gin.Context) {
	code := http.StatusNotFound
	message := "Item not found in Database"
	SendCodeWithMessage(c, code, message)
}

func SendDbErr(c *gin.Context, errMessage string) {
	code := http.StatusInternalServerError
	message := "Failed from database: " + errMessage
	SendCodeWithMessage(c, code, message)
}

func SendBadRequestErr(c *gin.Context, errMessage string) {
	code := http.StatusBadRequest
	message := "BadRequest: " + errMessage
	SendCodeWithMessage(c, code, message)
}

func SendInternalErr(c *gin.Context, errMessage string) {
	code := http.StatusInternalServerError
	message := "Internal Server Error: " + errMessage
	SendCodeWithMessage(c, code, message)
}

func SendUnauthorizedErr(c *gin.Context) {
	code := http.StatusUnauthorized
	message := "Unauthorized Option"
	SendCodeWithMessage(c, code, message)
}
