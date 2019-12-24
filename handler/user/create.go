package user

import (
	"fmt"
	"net/http"

	"miMallDemo/errno"
	"miMallDemo/handler"
	"miMallDemo/logger"

	"github.com/gin-gonic/gin"
)

// create a new user account
func Create(c *gin.Context)  {

	var r CreateRequest

	// Bind()：检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。
	// APIServer 采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的。
	if err := c.Bind(&r); err != nil {
		//  gin.H 简化生成 json 的方式
		// c.JSON(http.StatusOK, gin.H{ "status": "登录成功"})
		// 等价于
		// c.JSON(http.StatusOK, map[string]interface{}{ "status": "登录成功"})
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	admin2 := c.Param("username") // Param()：返回URL的参数值  这里对应的是router里面配置的relativePath"/:username"
	logger.Infof("URL username: %s", admin2)

	desc := c.Query("desc") // Query()：查询请求URL后面的参数，比如http://127.0.0.1:8080/v1/user/admin2?desc=test里面的desc
	logger.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type") // GetHeader():
	logger.Infof("Header Content-Type: %s", contentType)

	logger.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		handler.SendResponse(
			c,
			errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found")),
			nil,
		)
		return
	}

	if r.Password == "" {
		handler.SendResponse(
			c,
			fmt.Errorf("password is empty"),
			nil,
			)
		return
	}

	rsp := CreateResponse{Username: r.Username}

	handler.SendResponse(c, nil, rsp)

}

// 测试
// curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/admin2?desc=test -d'{"username":"admin","password":"admin"}'