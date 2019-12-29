package user

import (
	. "miMallDemo/handler"
	"miMallDemo/model"
	"miMallDemo/pkg/auth"
	"miMallDemo/pkg/errno"
	"miMallDemo/pkg/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.User
	// Req： curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/login -d'{"username":"admin","password":"admin"}'
	// Content-Type 为 application/json，表示消息体是json格式，且key->value: username->admin  password->admin
	//  Bind()：检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。
	// 示例代码，使用Bind()解析并绑定。
	// 结构体变量，Username值为admin，Password值为admin
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	} // 解析请求并绑定到结构体User变量u中，使用&u指向内存地址；如果出错，返回ErrBind

	d, err := model.GetUser(u.Username) // 按照用户名Username，从数据库中查询是否存在用户
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	} // 如果出错，返回ErrUserNotFound

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	} // 验证登录密码，如果密码不匹配，返回ErrPasswordIncorrect

	// 结构体变量的ID为d.ID;UserName为d.Username，d为数据库查询结果
	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "") // 对该用户进行签名，默认secret为空
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	} // 如果出错，返回ErrToken

	// 如果没有ErrBind、ErrUserNotFound、ErrPasswordIncorrect、ErrToken
	// 则将我们生成的token放入返回报文中返回
	SendResponse(c, nil, model.Token{Token: t})

}
