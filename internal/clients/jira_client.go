package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	jiraToken  string
	domainUrl  string
}

func NewJiraClient(timeout time.Duration, domainUrl, email, jiraToken string) JiraClient {
	return JiraClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		email:     email,
		jiraToken: jiraToken,
		domainUrl: domainUrl,
	}
}

func (c *JiraClient) GetUser() (JiraUser, error) {
	url := c.domainUrl + "rest/api/3/myself"
	fmt.Println("url: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return JiraUser{}, err
	}

	fmt.Println("email: ", c.email)
	fmt.Println("jiraTOken: ", c.jiraToken)

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.email, c.jiraToken)

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
