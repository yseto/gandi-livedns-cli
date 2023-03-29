package gandi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) GetSnapshots(domain string) ([]Snapshot, error) {
	url := c.BaseURL + "domains/" + domain + "/snapshots"

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

	var data []Snapshot
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetSnapshot(domain, id string) ([]Record, error) {
	url := c.BaseURL + "domains/" + domain + "/snapshots/" + id

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

	var data Snapshot
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data.ZoneData, nil
}

func (c *Client) DeleteSnapshot(domain, id string) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/snapshots/" + id

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		success := "Record was deleted"
		return &success, nil
	}
	return nil, ShowError(resp.Body)
}

func (c *Client) CreateSnapshot(domain, name string) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/snapshots"

	v, err := json.Marshal(CreateSnapshot{Name: name})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(v))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, ShowError(resp.Body)
	}

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data.Message, nil
}

func (c *Client) UpdateSnapshot(domain, id, name string) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/snapshots/" + id

	v, err := json.Marshal(CreateSnapshot{Name: name})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(v))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, ShowError(resp.Body)
	}

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data.Message, nil
}
