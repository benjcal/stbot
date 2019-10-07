package poloniex

import (
	"net/http"
)

type Client struct {
	key    string
	secret string
	client *http.Client
}

func NewClient(key, secret string) (c *Client) {
	return &Client{key: key, secret: secret, client: &http.Client{}}
}
