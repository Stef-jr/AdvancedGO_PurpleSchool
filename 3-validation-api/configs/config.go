package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email EmailConf
	Password PasswordConf
	Address AddressConf
}

type EmailConf struct {
	Email string
}

type PasswordConf struct {
	Password string
}

type AddressConf struct {
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Email: EmailConf{
			Email: os.Getenv("EMAIL"),
		},
		Password: PasswordConf{
			Password: os.Getenv("PASSWORD"),
		},
		Address: AddressConf{
			Address: os.Getenv("ADDRESS"),
		},
	}
}