package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var mutex = &sync.Mutex{}
var envConfig *EnvConfig

type EnvConfig struct {
	SERVER_PORT string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	// ABSOLUTE_PATH string
}

func initConfig() {
	log.Println("<- init config")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file, err:", err)
		return
	}

	log.Println("-> init config")
}

func GetConfig() *EnvConfig {
	log.Println("<- get config")

	mutex.Lock()
	defer mutex.Unlock()

	if envConfig == nil {
		initConfig()
	}

	var resConfig EnvConfig
	resConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	resConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	resConfig.DB_HOST = os.Getenv("DB_HOST")
	resConfig.DB_PORT = os.Getenv("DB_PORT")
	resConfig.DB_NAME = os.Getenv("DB_NAME")
	resConfig.SERVER_PORT = os.Getenv("SERVER_PORT")
	// resConfig.ABSOLUTE_PATH = os.Getenv("ABSOLUTE_PATH")

	log.Println("-> get config")
	return &resConfig
}
