package utils

import (
	"errors"
	"os"

	dotenv "github.com/joho/godotenv"
)

func GetEnv() (string, string, string, string, error) {
	err := dotenv.Load()
	if err != nil {
		return "", "", "", "", errors.New("Error loading .env file")
	}

	domainUrl := os.Getenv("DOMAIN_URL")
	email := os.Getenv("USER_EMAIL")
	jiraToken := os.Getenv("JIRA_API_TOKEN")
	tempoToken := os.Getenv("TEMPO_API_TOKEN")

	if domainUrl == "" || email == "" || jiraToken == "" || tempoToken == "" {
		return "", "", "", "", errors.New("You have an incomplete .env file.")
	}
	return domainUrl, email, jiraToken, tempoToken, nil
}
