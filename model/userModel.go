package model

import "time"

type userModel struct {
	id int
	userName string
	password string
	email string
	phone string
	role string
	createTime time.Time
	updateTime time.Time
}