package database

import (
	"fmt"
	"golang-united-lectures/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable TimeZone=Europe/Moscow", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil

}
