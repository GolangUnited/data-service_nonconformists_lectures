package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(host, port, user, password, name string) error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable TimeZone=Europe/Moscow", host, port, user, password, name)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil

}
