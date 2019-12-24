package handler

import (

	"net/http"

	"miMallDemo/errno"

	"github.com/gin-gonic/gin"

)

// 返回结构体 Code Message 通过 DecodeErr 解析而来
// Data interface{}类型 根据业务自己的需求来返回
type Response struct {
	Code int			`json:"code"`
	Message string		`json:"message"`
	Data interface{}	`json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{})  {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

