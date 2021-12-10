package marketo

type SmartCampaign struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateSmartCampaign(input SmartCampaign) (*SmartCampaign, error) {
	var result SmartCampaign
	return &result, nil
}

func (c *Client) GetSmartCampaign(id string) (*SmartCampaign, error) {
	var result SmartCampaign
	return &result, nil
}

func (c *Client) UpdateSmartCampaign(id string, input SmartCampaign) (*SmartCampaign, error) {
	var result SmartCampaign
	return &result, nil
}

func (c *Client) DeleteSmartCampaign(id string) error {
	return nil
}
