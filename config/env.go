package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "80 80"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "zr183114"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DBAddress", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, falback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return falback
}
