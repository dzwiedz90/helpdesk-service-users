package config

import (
	"log"
)

type Config struct {
	GRPCADDRESS string
	GRPCPORT    string
	TIMEOUT     string
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
}
