package router

import (

	"net/http"

	"miMallDemo/handler/sd"
	"miMallDemo/handler/user"
	"miMallDemo/router/middleware"

	"github.com/gin-gonic/gin"

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
	g.Use(middleware.Logger)
	g.Use(middleware.RequestId())
	g.Use(mw...)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect api router")
	})

	// http req and res demo
	//g.GET("/user/:id", func(c *gin.Context) {
	//	// Param()：返回URL的参数值
	//	c.Param("id") // id == "nicky"
	//
	//	// GET /path?id=1234&name=Manu&value=
	//	// Query()：查询请求URL后面的参数
	//	c.Query("id") == "1234"
	//	c.Query("name") == "Manu"
	//	c.Query("value") == ""
	//	c.Query("wtf") == ""
	//
	//	// DefaultQuery()：类似 Query()，但是如果 key 不存在，会返回默认值 查询请求URL后面的参数，如果没有填写默认值
	//	// GET /?name=Manu&lastname=
	//	c.DefaultQuery("name", "unknown") == "Manu"
	//	c.DefaultQuery("id", "none") == "none"
	//	c.DefaultQuery("lastname", "none") == ""
	//  c.PostForm("name") //从表单中查询参数
	//
	//	// Bind()：检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。
	//	// GetHeader()：获取 HTTP 头
	//})



	// 用户路由设置
	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)	// 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update) // 更新用户
		u.GET("", user.List) // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
		//u.POST("/:username", user.Create)
	}

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