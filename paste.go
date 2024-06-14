package pastebin

import (
	"net/url"
)

func (client *Client) CreatePaste(content string) (string, error) {
	pasteUrl, err := client.fetch("POST", "/api/api_post.php", &url.Values{
		"api_option":        {"paste"},
		"api_paste_private": {"2"},
		"api_paste_title":   {"Created with simonkarman/pastebin-client-go"},
		"api_paste_code":    {content},
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

func (client *Client) GetPaste(pasteKey string) (string, error) {
	pasteContent, err := client.fetch("POST", "/api/api_raw.php", &url.Values{
		"api_option":    {"show_paste"},
		"api_paste_key": {pasteKey},
	})
	return pasteContent, err
}

func (client *Client) DeletePaste(pasteKey string) error {
	_, err := client.fetch("POST", "/api/api_post.php", &url.Values{
		"api_option":    {"delete"},
		"api_paste_key": {pasteKey},
	})
	return err
}
