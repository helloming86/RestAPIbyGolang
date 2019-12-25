package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GetShortId() (string, error) {
	return shortid.Generate()
}

func GenReqId(c *gin.Context) string {
	// 使用 requestId 标记每次请求全链路日志
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}