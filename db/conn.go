package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Song{})
	return db
}