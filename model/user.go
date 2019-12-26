package model

import (
	"fmt"

	"miMallDemo/auth"

	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

// 数据库表
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "tb_users"
}

func (u *User) Create() error {
	return DB.Create(&u).Error
}

func DeleteUser(id uint) error {
	user := User{}
	user.ID = id
	return DB.Delete(&user).Error
}

func (u *User) Update() error {
	return DB.Save(u).Error
}

func GetUser(username string) (*User, error) {
	u := &User{}
	d := DB.Where("username = ?", username).First(&u)
	return u, d.Error
}

func ListUser(username string, offset, limit int) ([]*User, uint, error) {
	if limit == 0 {
		limit = 50
	}

	users := make([]*User, 0)
	var count uint

	where := fmt.Sprintf("username like '%%%s%%'", username)

	if err := DB.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil

}

// 关于该函数的定义
// 1. 函数签名 返回 直接声明了一个变量 err error
// 2. 函数体 err = 而不是 err :=
// 3. 函数体 return 后面没有err
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
