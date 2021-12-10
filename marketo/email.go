package marketo

type Email struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateEmail(program Email) (*Email, error) {
	var result Email
	return &result, nil
}

func (c *Client) GetEmail(id string) (*Email, error) {
	var result Email
	return &result, nil
}

func (c *Client) UpdateEmail(id string, program Email) (*Email, error) {
	var result Email
	return &result, nil
}

func (c *Client) DeleteEmail(id string) error {
	return nil
}
