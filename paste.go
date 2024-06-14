package pastebin

import (
	"net/url"
)

func (client *Client) CreatePaste(title string, format string, content string) (string, error) {
	pasteUrl, err := client.fetch("POST", "/api/api_post.php", &url.Values{
		"api_option":       {"paste"},
		"api_paste_name":   {title},
		"api_paste_format": {format},
		"api_paste_code":   {content},
	})
	if err != nil {
		return "", err
	}

	pasteUrl_, err := url.Parse(pasteUrl)
	if err != nil {
		return "", err
	}
	pasteKey := pasteUrl_.Path
	if pasteKey[0] == '/' {
		pasteKey = pasteKey[1:]
	}
	return pasteKey, nil
}

func (client *Client) DeletePaste(pasteKey string) error {
	_, err := client.fetch("POST", "/api/api_post.php", &url.Values{
		"api_option":    {"delete"},
		"api_paste_key": {pasteKey},
	})
	return err
}
