package irstats

import "net/http"

type CarsDriven = []int

// CarsDriven returns which cars (car_id) someone has driven
func (c *Client) CarsDriven(custID *string) (*CarsDriven, *http.Response, error) {
	p := map[string]string{
		"custid": *custID,
	}

	carsDriven := &CarsDriven{}
	resp, err := c.do(UrlPathCarsDriven, &p, carsDriven)
	return carsDriven, resp, err
}
