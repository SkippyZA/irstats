package irstats

import (
	"net/http"
	"net/url"
)

type CarsDriven = []int

// CarsDriven returns which cars (car_id) someone has driven
func (c *Client) CarsDriven(custID *string) (*CarsDriven, *http.Response, error) {
	v := url.Values{"custid": {*custID}}
	carsDriven := &CarsDriven{}
	resp, err := c.do(URLPathCarsDriven, &v, carsDriven)
	return carsDriven, resp, err
}
