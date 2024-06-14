package pastebin

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Host    url.URL
	DevKey  string
	UserKey string
}

func New(host url.URL, devKey string, userKey string) *Client {
	return &Client{
		Host:    host,
		DevKey:  devKey,
		UserKey: userKey,
	}
}

func (client *Client) fetch(httpMethod string, path string, data *url.Values) (string, error) {
	// Create URL
	relativePath, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	httpUrl := client.Host.ResolveReference(relativePath)

	// Create Body
	if data == nil {
		data = &url.Values{"api_dev_key": {client.DevKey}}
		data = &url.Values{"api_user_key": {client.UserKey}}
	} else {
		data.Set("api_dev_key", client.DevKey)
		data.Set("api_user_key", client.UserKey)
	}
	httpBody := data.Encode()

	// Create the fetch
	req, err := http.NewRequest(httpMethod, httpUrl.String(), strings.NewReader(httpBody))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// Return the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("%s failed [%d] %s", httpMethod, resp.StatusCode, body)
	}
	return string(body), nil
}
