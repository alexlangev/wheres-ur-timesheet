package clients

import (
	"net/http"
	"time"
	"net/http"
)

const baseUrl = "TODO/rest/api/3/"

type JiraIssue struct {
	Id     string        `json:"id"`
	Key    string        `json:"key"`
}

type JiraUser struct {
	email string
	accountId string
}

type JiraClient struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) JiraClient {
	return JiraClient {
		httpClient: http.Client{
			Timeout: timeout,
		}
	}
}

// get current user

// read from env email, token
func (c *Client) GetUser() (JiraUser, error) {
	url := baseUrl + "myself"
	// email from env
	// jira token from env
	
	// read from cache?

	// make call

	// unmarshal

	// add to cache
}


// client methods
// get issue based on key
