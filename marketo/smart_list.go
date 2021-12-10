package marketo

type SmartList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateSmartList(program SmartList) (*SmartList, error) {
	var result SmartList
	return &result, nil
}

func (c *Client) GetSmartList(id string) (*SmartList, error) {
	var result SmartList
	return &result, nil
}

func (c *Client) UpdateSmartList(id string, program SmartList) (*SmartList, error) {
	var result SmartList
	return &result, nil
}

func (c *Client) DeleteSmartList(id string) error {
	return nil
}
