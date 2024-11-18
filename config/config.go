package config

import (
	"os"
)

type Config struct {
	Port         string
	Host         string
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewConfig() Config {
	return Config{
		Port:         "8080",
		Host:         os.Getenv("HOST"),
		SMTPServer:   os.Getenv("SMTP_SERVER"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
