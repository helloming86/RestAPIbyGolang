package user

import (
	"fmt"
	"net/http"

	"miMallDemo/errno"
	"miMallDemo/logger"

	"github.com/gin-gonic/gin"

)

// create a new user account
func Create(c *gin.Context)  {

	var r struct{
		Username string `json:"username"`
		Password string	`json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	logger.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		logger.Errorf("UserErrrrrr", "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		logger.Debug("err type is ErrUserNotFound")
	}
	
	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

}
