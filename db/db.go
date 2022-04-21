package db

import (
	"fiber_demo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/go_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
