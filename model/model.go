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
type UserList struct {
	Lock *sync.Mutex
	IdMap map[uint]*UserInfo
}

// 非数据库表
type Token struct {
	Token string			`json:"token"`
}
