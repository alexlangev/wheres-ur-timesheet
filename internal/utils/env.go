package utils

import (
	"errors"
	"os"

	dotenv "github.com/joho/godotenv"
)

func ValidateEnv() error {
	err := dotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}

	domainUrl := os.Getenv("DOMAIN_URL")
	email := os.Getenv("USER_EMAIL")
	jiraToken := os.Getenv("JIRA_API_TOKEN")
	tempoToken := os.Getenv("TEMPO_API_TOKEN")

	if domainUrl == "" || email == "" || jiraToken == "" || tempoToken == "" {
		return errors.New("❌ incomplete .env file")
	}
	return nil
}
