package gandi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetDomains() ([]Domain, error) {
	url := c.BaseURL + "domains"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, ShowError(resp.Body)
	}

	var data []Domain
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
