package router

import (
	"github.com/gin-gonic/gin"
	"miMallDemo/handler/sd"
	"miMallDemo/router/middleware"
	"net/http"
)


// Load: loads the middlewares, routes, handlers.
// func test1(args ...string) 可以接受任意个string参数
// test1(strss...) 切片被打散传入
// strss=append(strss,strss2...) strss2的元素被打散一个个append进strss

// gin.Recovery()：	在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic
//					这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
// middleware.NoCache：强制浏览器不使用缓存
// middleware.Options：浏览器跨域 OPTIONS 请求设置
// middleware.Secure：一些安全设置


func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect api router")
	})

	// the health check handler
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}