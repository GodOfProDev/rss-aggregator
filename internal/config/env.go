package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Env struct {
	Port int
}

var env *Env

func GetEnv() *Env {
	if env != nil {
		return env
	}

	return &Env{Port: 8080}
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("PORT")

	GetEnv().Port, err = strconv.Atoi(portString)

	return err
}
