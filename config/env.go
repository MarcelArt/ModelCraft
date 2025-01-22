package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	PORT         string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBHost       string
	JwtSecret    string
	ServerENV    string
	SMTPHost     string
	SMTPPort     int
	SMTPName     string
	SMTPEmail    string
	SMTPPassword string
}

var Env *env

func SetupENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Panic(err.Error())
	}

	Env = &env{
		PORT:         os.Getenv("PORT"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBHost:       os.Getenv("DB_HOST"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		ServerENV:    os.Getenv("SERVER_ENV"),
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     smtpPort,
		SMTPName:     os.Getenv("SMTP_NAME"),
		SMTPEmail:    os.Getenv("SMTP_EMAIL"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
