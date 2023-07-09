package config

import (
	"docker_go_test/app/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if there is an error opening the connection, handle it
	exception.PanicIfNeeded(err)

	return db
}
