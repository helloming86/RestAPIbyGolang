package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId int64
	UserName string
}
