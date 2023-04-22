package kagi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ClientConfig struct {
	authToken string

	BaseURL string
}

type Client struct {
	config ClientConfig
}

func NewClient(token string) *Client {
	return &Client{
		config: ClientConfig{
			authToken: token,
			BaseURL:   "https://kagi.com/api/v0",
		},
	}
}

func (c *Client) apiURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) sendRequest(req *http.Request, v any) (err error) {
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.config.authToken))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	if v == nil {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(&v)
	return
}

func (c *Client) buildRequest(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	)
}
