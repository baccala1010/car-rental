package config

import "os"

type Config struct {
	// db config
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// jwt config
	JWTSecret string

	// smtp config
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string

	AdminRegistrationSecret string
}

func LoadConfig() Config {
	return Config{
		DBHost:                  os.Getenv("DB_HOST"),
		DBPort:                  os.Getenv("DB_PORT"),
		DBUser:                  os.Getenv("DB_USER"),
		DBPassword:              os.Getenv("DB_PASSWORD"),
		DBName:                  os.Getenv("DB_NAME"),
		JWTSecret:               os.Getenv("JWT_SECRET"),
		SMTPHost:                os.Getenv("SMTP_HOST"),
		SMTPPort:                os.Getenv("SMTP_PORT"),
		SMTPUser:                os.Getenv("SMTP_USER"),
		SMTPPassword:            os.Getenv("SMTP_PASSWORD"),
		SMTPFrom:                os.Getenv("SMTP_FROM"),
		AdminRegistrationSecret: os.Getenv("ADMIN_REGISTRATION_SECRET"),
	}
}
