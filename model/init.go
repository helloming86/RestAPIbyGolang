package model

import (
	"miMallDemo/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
	DB = openDB()
}

func openDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		logger.Info("failed to connect database")
	}
	logger.Info("DataBase is Ready")

	// set for db conn
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(true)           //
	db.DB().SetMaxOpenConns(0) // 用于设置最大打开的连接数，默认值为0表示不限制
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}
