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
