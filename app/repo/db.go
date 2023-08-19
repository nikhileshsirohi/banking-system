package repo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() error {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/banking_system?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB_INIT_ERROR", err)
	}
	return err
}
