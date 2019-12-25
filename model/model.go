package model

import "sync"


// 非数据库表
type UserInfo struct {
	ID        	uint		`json:"id"`
	Username 	string		`json:"username"`
	SayHello 	string		`json:"sayHello"`
	Password 	string		`json:"password"`
	CreatedAt 	string		`json:"createdAt"`
	UpdatedAt 	string		`json:"updatedAt"`
}

// 非数据库表
// 将UserInfo封装为一个带锁的新的结构体
// 实际业务，SayHello 包含shortID
// 多条数据，加同步锁，处理shortID
type UserList struct {
	Lock *sync.Mutex
	IdMap map[uint]*UserInfo
}

// 非数据库表
type Token struct {
	Token string			`json:"token"`
}
