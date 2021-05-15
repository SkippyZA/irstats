package irstats

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	resty "github.com/go-resty/resty/v2"
)

const (
	defaultUserAgent = "go-irstats"
	defaultBaseURL   = "https://members.iracing.com"
)

var (
	ErrAuthenticationFailed = errors.New("Failed to authenticate with iRacing")
)

type Client struct {
	http *resty.Client // http client

	username      string
	password      string
	userAgent     string
	authenticated bool
}

func NewClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	http := resty.New()
	http.SetHostURL(defaultBaseURL)
	http.SetRedirectPolicy(resty.NoRedirectPolicy())

	c := &Client{
		http:      http,
		username:  username,
		password:  password,
		userAgent: defaultUserAgent,
	}

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) do(path urlPath, params *map[string]string, v interface{}) (*http.Response, error) {
	if err := c.assertLoggedIn(); err != nil {
		return nil, err
	}

	var resp *resty.Response
	var err error

	r := c.http.R().SetHeader("User-Agent", c.username)
	if params == nil {
		resp, err = r.Get(path)
	} else {
		resp, err = r.SetFormData(*params).Post(path)
	}
	rawResponse := resp.RawResponse

	if err != nil {
		log.Printf("API request %s failed. %v\n", path, err)
		return rawResponse, err
	}

	// If there is no target to unmarshal the json to, then we return out here
	if v == nil {
		return rawResponse, nil
	}

	err = json.Unmarshal(resp.Body(), v)
	if err != nil {
		log.Println("Failed to unmarshal response body", err)
		return rawResponse, err
	}

	return rawResponse, nil
}

func (c *Client) login() error {
	// TODO: need to add in some locking so that we dont spam the auth
	resp, _ := c.http.R().
		SetFormData(map[string]string{
			"username": c.username,
			"password": c.password,
		}).
		Post(URLPathLogin)

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "irsso_membersv2" && len(cookie.Value) > 0 {
			c.authenticated = true
			return nil
		}
	}

	c.authenticated = false
	return ErrAuthenticationFailed
}

func (c *Client) assertLoggedIn() error {
	if c.authenticated {
		return nil
	}

	return c.login()
}
