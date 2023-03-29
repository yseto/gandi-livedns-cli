package gandi

import "net/http"

type Client struct {
	BaseURL string
	Client  *http.Client
}

type AuthTransport struct {
	T      http.RoundTripper
	apiKey string
}

func (adt *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Apikey "+adt.apiKey)
	return adt.T.RoundTrip(req)
}

func NewClient(apikey string) *Client {
	c := http.DefaultClient
	c.Transport = &AuthTransport{T: http.DefaultTransport, apiKey: apikey}

	return &Client{
		BaseURL: "https://api.gandi.net/v5/livedns/",
		Client:  c,
	}
}
