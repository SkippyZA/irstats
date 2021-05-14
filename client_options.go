package irstats

// ClientOptionFunc can be used customize a new iRacing Stats API client.
type ClientOptionFunc func(*Client) error

// WithBaseURL sets the base URL for API requests to a custom endpoint.
func WithBaseURL(s string) ClientOptionFunc {
	return func(c *Client) error {
		c.http.SetHostURL(s)
		return nil
	}
}

// WithUserAgent set the user agent to be used for API requests
func WithUserAgent(s string) ClientOptionFunc {
	return func(c *Client) error {
		c.userAgent = s
		return nil
	}
}
