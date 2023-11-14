package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// DB global var for connect DB
var DB *gorm.DB

type DBConfig struct {
	user     string
	password string
	host     string
	name     string
	port     string
}

func BuildDBConfig() *DBConfig {
	// Load environment variables from the .env file
	godotenv.Load()

	dbConfig := DBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		password: os.Getenv("DB_PASSWORD"),
		user:     os.Getenv("DB_USER"),
		name:     os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.name)
}
