package jiraapi

import (
	"net/http"
)

type Client struct {
	httpClient http.Client
}
