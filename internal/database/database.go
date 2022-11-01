package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	Client *gorm.DB
}

var db database

func New(host, port, user, password, name string) (*database, error) {

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " database=" + name + " sslmode=disable TimeZone=Europe/Moscow"

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db = database{
		Client: client,
	}

	return &db, nil

}

func GetInstance() *database {
	return &db
}

func (d *database) AutoMigrate(i ...interface{}) error {
	return d.Client.AutoMigrate(i...)
}

func (d *database) Close() error {

	db, err := d.Client.DB()

	if err != nil {
		return err
	}

	db.Close()

	return nil

}
