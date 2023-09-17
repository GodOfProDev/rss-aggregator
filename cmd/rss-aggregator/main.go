package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
