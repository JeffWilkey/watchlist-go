package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBName string
var MongoURI string
var JWTSecret string

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBName = os.Getenv("DB_NAME")
	MongoURI = os.Getenv("MONGO_URI")
	JWTSecret = os.Getenv("JWT_SECRET")
}