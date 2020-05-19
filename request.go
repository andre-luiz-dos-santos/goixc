package goixc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	*Client
	Form   string
	Body   interface{}
	Listar bool
}

func (c *Client) Get(form string, body interface{}) *Request {
	return &Request{
		Client: c,
		Form:   form,
		Body:   body,
	}
}

func (c *Client) List(form string, body interface{}) *Request {
	return &Request{
		Client: c,
		Form:   form,
		Body:   body,
		Listar: true,
	}
}

func (c *Request) Run(ctx context.Context) ([]byte, error) {
	reqBody, err := c.bodyReader()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.url(), reqBody)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	if c.Listar {
		req.Header.Set("ixcsoft", "listar")
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *Request) RunJSON(ctx context.Context, respJSON interface{}) error {
	respBody, err := c.Run(ctx)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, respJSON)
	if err != nil {
		return &ErrBadJSON{respBody, err}
	}
	return nil
}

func (c *Request) url() string {
	return c.formURL(c.Form)
}

func (c *Request) bodyReader() (io.Reader, error) {
	b, ok := c.Body.([]byte)
	if ok {
		return bytes.NewReader(b), nil
	}
	b, err := json.Marshal(c.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
