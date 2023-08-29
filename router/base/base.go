package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess      = 0
	CodeFail         = -1
	CodeUnauthorized = 403
)

type Response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data:    data,
		Code:    CodeSuccess,
		Message: "",
	})
	return
}

func SuccessMsg(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data:    data,
		Code:    CodeSuccess,
		Message: message,
	})
	return
}

func ReturnDataDirect(c *gin.Context, data interface{}) {
	c.PureJSON(http.StatusOK, data)
	return
}

func Fail(c *gin.Context, err string) {
	c.JSON(http.StatusOK, Response{
		Code:    -1,
		Message: err,
	})
	return
}
