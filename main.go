package main

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	email := os.Getenv("USER_EMAIL")

	fmt.Println("Where's your timesheet " + email + "?")

}
