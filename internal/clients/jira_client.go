package clients

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type JiraIssue struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

type JiraUser struct {
	Email       string `json:"emailAddress"`
	AccountId   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

type JiraClient struct {
	httpClient http.Client
	email      string
	apiToken   string
	baseUrl    string
}

func NewClient(timeout time.Duration) JiraClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	email := os.Getenv("USER_EMAIL")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	baseUrl := fmt.Sprintf("%s/rest/api/3/", os.Getenv("DOMAIN_URL"))

	if email == "" || apiToken == "" {
		log.Fatal("❌ Missing USER_EMAIL or JIRA_API_TOKEN in environment")
	}

	return JiraClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		email:    email,
		apiToken: apiToken,
		baseUrl:  baseUrl,
	}
}

func (c *JiraClient) GetUser() (JiraUser, error) {
	url := c.baseUrl + "myself"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return JiraUser{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.email, c.apiToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return JiraUser{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return JiraUser{}, fmt.Errorf("unexpected response %d: %s", resp.StatusCode, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return JiraUser{}, err
	}

	currentUser := JiraUser{}
	err = json.Unmarshal(data, &currentUser)
	if err != nil {
		return JiraUser{}, err
	}

	return currentUser, nil
}
