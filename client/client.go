package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) Ping() (map[string]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/ping", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
