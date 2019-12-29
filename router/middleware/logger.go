package middleware

import (
	"github.com/gin-gonic/gin"
	"miMallDemo/pkg/logger"
)

// 将zap logger 作为 中间件使用
func Logger(c *gin.Context) {
	c.Next()
	logger.Infof("Incoming Request: %s, status: %d ", c.Request.URL.Path, c.Writer.Status())
}
