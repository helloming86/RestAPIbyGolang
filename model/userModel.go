package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id int64
	UserId int64
	AddId int64
	Name string
	Address string
	Email string
	Phone string
}

type UserAuth struct {
	gorm.Model
	ProfileID int      `gorm:"not null" json:"profile_id"`
	Username  string   `gorm:"size:20;unique_index" json:"username"`
	Email     string   `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string   `gorm:"not null" json:"password"`
	Remember  bool     `gorm:"not null" json:"remeber"`
	TwoFA     bool     `gorm:"not null" json:"twofa"`
	Access    int      `gorm:"not null" json:"access"`
	State     int      `gorm:"not null" json:"state"`
	LastSeen  string   `gorm:"not null" json:"lastseen"`
}
