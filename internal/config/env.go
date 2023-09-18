package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
)

type EnvConfig struct {
	Port  int
	DbUrl string
}

var envInstance *EnvConfig
var envOnce sync.Once

func GetEnv() *EnvConfig {
	envOnce.Do(func() {
		envInstance = &EnvConfig{Port: 8080}
	})

	return envInstance
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("PORT")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return errors.New("DB_URL was not found in the environment")
	}

	GetEnv().Port, err = strconv.Atoi(portString)
	GetEnv().DbUrl = dbUrl

	return err
}
