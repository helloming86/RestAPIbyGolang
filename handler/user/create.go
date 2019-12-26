package user

import (
	"miMallDemo/errno"
	. "miMallDemo/handler"
	"miMallDemo/logger"
	"miMallDemo/model"
	"miMallDemo/utils"

	"github.com/gin-gonic/gin"
)

// 创建用户逻辑
// 		1. 从 HTTP 消息体获取参数（用户名和密码）
//		2. 参数校验
//		3. 加密密码
//		4. 在数据库中添加数据记录
//		5. 返回结果（这里是用户名）

func Create(c *gin.Context) {

	logger.Infof("User Create function called. X-Request-Id is %d", utils.GenReqId(c))

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{Username: r.Username}

	SendResponse(c, nil, rsp)

}
