package marketo

import (
	"net/http"
	"time"
)

type Client struct {
	ID         string
	Secret     string
	URL        string
	HTTPClient *http.Client
}

func NewClient(url string, id string, secret string) (*Client, error) {
	return &Client{
		URL:        url,
		ID:         id,
		Secret:     secret,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (c *Client) CreateProgram(program Program) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) GetProgram(id string) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) UpdateProgram(id string, program Program) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) DeleteProgram(id string) error {
	return nil
}
