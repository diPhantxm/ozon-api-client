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

	analytics     *Analytics
	fbo           *FBO
	fbs           *FBS
	finance       *Finance
	products      *Products
	promotions    *Promotions
	rating        *Rating
	warehouses    *Warehouses
	returns       *Returns
	reports       *Reports
	cancellations *Cancellations
	categories    *Categories
	polygons      *Polygons
}

func (c Client) Analytics() *Analytics {
	return c.analytics
}

func (c Client) FBO() *FBO {
	return c.fbo
}

func (c Client) FBS() *FBS {
	return c.fbs
}

func (c Client) Finance() *Finance {
	return c.finance
}

func (c Client) Products() *Products {
	return c.products
}

func (c Client) Promotions() *Promotions {
	return c.promotions
}

func (c Client) Rating() *Rating {
	return c.rating
}

func (c Client) Warehouses() *Warehouses {
	return c.warehouses
}

func (c Client) Returns() *Returns {
	return c.returns
}

func (c Client) Reports() *Reports {
	return c.reports
}

func (c Client) Cancellations() *Cancellations {
	return c.cancellations
}

func (c Client) Categories() *Categories {
	return c.categories
}

func (c Client) Polygons() *Polygons {
	return c.polygons
}

func NewClient(clientId, apiKey string) *Client {
	coreClient := core.NewClient(DefaultAPIBaseUrl, map[string]string{
		"Client-Id": clientId,
		"Api-Key":   apiKey,
	})

	return &Client{
		client:        coreClient,
		analytics:     &Analytics{client: coreClient},
		fbo:           &FBO{client: coreClient},
		fbs:           &FBS{client: coreClient},
		finance:       &Finance{client: coreClient},
		products:      &Products{client: coreClient},
		promotions:    &Promotions{client: coreClient},
		rating:        &Rating{client: coreClient},
		warehouses:    &Warehouses{client: coreClient},
		returns:       &Returns{client: coreClient},
		reports:       &Reports{client: coreClient},
		cancellations: &Cancellations{client: coreClient},
		categories:    &Categories{client: coreClient},
		polygons:      &Polygons{client: coreClient},
	}
}

func NewMockClient(handler http.HandlerFunc) *Client {
	coreClient := core.NewMockClient(handler)

	return &Client{
		client:        coreClient,
		analytics:     &Analytics{client: coreClient},
		fbo:           &FBO{client: coreClient},
		fbs:           &FBS{client: coreClient},
		finance:       &Finance{client: coreClient},
		products:      &Products{client: coreClient},
		promotions:    &Promotions{client: coreClient},
		rating:        &Rating{client: coreClient},
		warehouses:    &Warehouses{client: coreClient},
		returns:       &Returns{client: coreClient},
		reports:       &Reports{client: coreClient},
		cancellations: &Cancellations{client: coreClient},
		categories:    &Categories{client: coreClient},
		polygons:      &Polygons{client: coreClient},
	}
}
