package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	var envs map[string]string
	envs, err := godotenv.Read("./config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var environmentalVariable = envs[name]
	return environmentalVariable
}
