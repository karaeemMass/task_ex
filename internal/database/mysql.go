package database

import (
	"fmt"
	"task_ex/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB() (*gorm.DB, error) {

	user := "root"
	password := ""
	host := "127.0.0.1:3306"
	dbname := "grpc_task"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the Task model
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
