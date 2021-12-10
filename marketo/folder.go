package marketo

type Folder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateFolder(program Folder) (*Folder, error) {
	var result Folder
	return &result, nil
}

func (c *Client) GetFolder(id string) (*Folder, error) {
	var result Folder
	return &result, nil
}

func (c *Client) UpdateFolder(id string, program Folder) (*Folder, error) {
	var result Folder
	return &result, nil
}

func (c *Client) DeleteFolder(id string) error {
	return nil
}
