package database

import (
	"fmt"
	"log"
	"task_ex/config"
	"task_ex/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Use the typed Database config returned by LoadConfig
	dbCfg := cfg.Database
	user := dbCfg.User
	password := dbCfg.Password
	host := dbCfg.Host
	port := dbCfg.Port
	dbname := dbCfg.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the models
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.RefreshToken{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.PricesGold{}); err != nil {
		return nil, err
	}

	return db, nil
}
