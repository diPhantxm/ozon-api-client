package ozon

import (
	"net/http"
	"testing"
)

const (
	apiKey   = "some_key"
	clientId = "some_client_id"
)

func TestNewClient(t *testing.T) {
	client := NewClient(
		WithAPIKey(apiKey),
		WithClientId(clientId),
		WithURI(DefaultAPIBaseUrl),
		WithHttpClient(http.DefaultClient),
	)

	if client.client.Options["Api-Key"] != apiKey {
		t.Errorf("expected api key: %s, but got: %s", apiKey, client.client.Options["Api-Key"])
	}
	if client.client.Options["Client-Id"] != clientId {
		t.Errorf("expected client id: %s, but got: %s", clientId, client.client.Options["Client-Id"])
	}
}
