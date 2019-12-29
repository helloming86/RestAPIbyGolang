package user

import (
	"miMallDemo/pkg/errno"
	"strconv"

	. "miMallDemo/handler"
	"miMallDemo/model"
	"miMallDemo/pkg/logger"
	"miMallDemo/utils"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	logger.Infof("User Create function called. X-Request-Id is %d", utils.GenReqId(c))

	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.User
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.ID = uint(userId)

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
