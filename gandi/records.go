package gandi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) GetRecords(domain string) ([]Record, error) {
	url := c.BaseURL + "domains/" + domain + "/records"

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

	var data []Record
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CreateRecord(domain string, record Record) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/records"

	v, err := json.Marshal(record)
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

type ReplaceRecords struct {
	Items []Record `json:"items"`
}

func (c *Client) ReplaceRecord(domain string, record ReplaceRecord) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/records/" + record.ZoneName

	v, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(v))
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

func (c *Client) DeleteRecord(domain string, record Record) (*string, error) {
	url := c.BaseURL + "domains/" + domain + "/records/" + record.Name + "/" + record.Type

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
