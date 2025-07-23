package config

import (
	"os"

	"github.com/joho/godotenv"
)

var REGISTRATION_CODE string
var DEFAULT_PASSWORD string
var SECRETKEY string
var HEADER string
var DOMAIN string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	REGISTRATION_CODE = os.Getenv("REGISTRATION_CODE")
	if REGISTRATION_CODE == "" {
		panic("REGISTRATION_CODE is not set")
	}
	DEFAULT_PASSWORD = os.Getenv("DEFAULT_PASSWORD")
	if DEFAULT_PASSWORD == "" {
		panic("DEFAULT_PASSWORD is not set")
	}
	SECRETKEY = os.Getenv("SECRETKEY")
	if SECRETKEY == "" {
		panic("SECRETKEY is not set")
	}
	HEADER = os.Getenv("HEADER")
	if HEADER == "" {
		panic("HEADER is not set")
	}
	DOMAIN = os.Getenv("DOMAIN")
	if DOMAIN == "" {
		panic("DOMAIN is not set")
	}
}
