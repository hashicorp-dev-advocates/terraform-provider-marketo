package marketo

type EmailTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateEmailTemplate(program EmailTemplate) (*EmailTemplate, error) {
	var result EmailTemplate
	return &result, nil
}

func (c *Client) GetEmailTemplate(id string) (*EmailTemplate, error) {
	var result EmailTemplate
	return &result, nil
}

func (c *Client) UpdateEmailTemplate(id string, program EmailTemplate) (*EmailTemplate, error) {
	var result EmailTemplate
	return &result, nil
}

func (c *Client) DeleteEmailTemplate(id string) error {
	return nil
}
