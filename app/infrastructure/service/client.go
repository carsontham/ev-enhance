package service

import (
	"ev-enhance/app/httpapi"
	"net/http"
)

var _ ExternalPartnerAPI = new(Client)

type Client struct {
	getEVRateEndpoint string
	httpClient        httpapi.Caller
}

// NewClient ...
func NewClient(getEVRateEndpoint string) *Client {
	return &Client{
		getEVRateEndpoint: getEVRateEndpoint,
		httpClient:        httpapi.NewCaller(http.DefaultClient),
	}
}
