package goixc

import (
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	URL      string
	Username string
	Password string

	http http.Client
}

func NewClient(url, token string) (*Client, error) {
	t := strings.SplitN(token, ":", 2)
	if len(t) != 2 {
		return nil, &ErrBadToken{token}
	}
	return &Client{
		URL:      url,
		Username: t[0],
		Password: t[1],
	}, nil
}

func (c *Client) formURL(form string) string {
	return fmt.Sprintf("%s/webservice/v1/%s", c.URL, form)
}
