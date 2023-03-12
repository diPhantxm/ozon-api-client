package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

const (
	DefaultAPIBaseUrl = "https://api-seller.ozon.ru"
)

type Client struct {
	client *core.Client
}

func NewClient(clientId, apiKey string) *Client {
	return &Client{
		client: core.NewClient(DefaultAPIBaseUrl, map[string]string{
			"Client-Id": clientId,
			"Api-Key":   apiKey,
		}),
	}
}

func NewMockClient(handler http.HandlerFunc) *Client {
	return &Client{
		client: core.NewMockClient(handler),
	}
}
