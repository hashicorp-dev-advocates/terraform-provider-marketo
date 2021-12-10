package marketo

type Program struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateProgram(input Program) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) GetProgram(id string) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) UpdateProgram(id string, input Program) (*Program, error) {
	var result Program
	return &result, nil
}

func (c *Client) DeleteProgram(id string) error {
	return nil
}
