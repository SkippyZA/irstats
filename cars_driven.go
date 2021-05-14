package irstats

type CarsDriven = []int

// CarsDriven Returns which cars (car_id) someone has driven
func (c *Client) CarsDriven(custID *string) (*CarsDriven, error) {
	p := map[string]string{
		"custid": *custID,
	}

	carsDriven := &CarsDriven{}
	err := c.do(UrlPathCarsDriven, &p, carsDriven)
	return carsDriven, err
}
