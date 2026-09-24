package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, errMessages ...string) {
	var errs interface{}
	if len(errMessages) > 0 {
		errs = errMessages
	} else {
		errs = []string{message}
	}

	c.JSON(code, Response{
		Status:  false,
		Message: message,
		Errors:  errs,
		Data:    nil,
	})
}
