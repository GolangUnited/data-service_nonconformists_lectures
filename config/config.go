package config

import (
	"errors"
	"os"
)

const PROTOCOL_TCP = "tcp"
const PORT_8080 = "8080"

var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_DATABASE string

func Get() error {

	DB_HOST = os.Getenv("LECTURES_DB_HOST")
	DB_PORT = os.Getenv("LECTURES_DB_PORT")
	DB_USER = os.Getenv("LECTURES_DB_USER")
	DB_PASSWORD = os.Getenv("LECTURES_DB_PASSWORD")
	DB_DATABASE = os.Getenv("LECTURES_DB_NAME")

	if DB_HOST == "" {
		return errors.New("env varialble lectures_db_host has not filled")
	}

	if DB_PORT == "" {
		return errors.New("env varialble lectures_db_port has not filled")
	}

	if DB_USER == "" {
		return errors.New("env varialble lectures_db_user has not filled")
	}

	if DB_PASSWORD == "" {
		return errors.New("env varialble lectures_db_password has not filled")
	}

	if DB_DATABASE == "" {
		return errors.New("env varialble lectures_db_name has not filled")
	}

	return nil

}
