package user


import (

	"strconv"

	"miMallDemo/errno"
	. "miMallDemo/handler"
	"miMallDemo/model"

	"github.com/gin-gonic/gin"

)

func Delete(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}