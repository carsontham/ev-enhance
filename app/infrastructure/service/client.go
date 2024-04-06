package service

import (
	"ev-enhance/app/httpapi"
	"net/http"
)

var _ ExternalPartnerAPI = new(Client)

type Client struct {
	translateEndpoint string
	httpClient        httpapi.Caller
}

// NewClient ...
func NewClient(translateEndpoint string) *Client {
	return &Client{
		translateEndpoint: translateEndpoint,
		httpClient:        httpapi.NewCaller(http.DefaultClient),
	}
}
