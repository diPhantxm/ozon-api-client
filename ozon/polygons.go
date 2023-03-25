package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Polygons struct {
	client *core.Client
}

type CreateDeliveryPolygonParams struct {
	// Delivery polygon coordinates in [[[lat long]]] format
	Coordinates string `json:"coordinates"`
}

type CreateDeliveryPolygonResponse struct {
	core.CommonResponse

	// Polygon identifier
	PolygonId int64 `json:"polygon_id"`
}

// You can link a polygon to the delivery method.
//
// Create a polygon getting its coordinates on https://geojson.io: mark at least 3 points on the map and connect them
func (c Polygons) CreateDelivery(params *CreateDeliveryPolygonParams) (*CreateDeliveryPolygonResponse, error) {
	url := "/v1/polygon/create"

	resp := &CreateDeliveryPolygonResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type LinkDeliveryMethodToPolygonParams struct {
	// Delivery method identifier
	DeliveryMethodId int32 `json:"delivery_method_id"`

	// Polygons list
	Polygons []LinkDeliveryMethodToPolygonPolygon `json:"polygons"`

	// Warehouse location
	WarehouseLocation LinkDeliveryMethodToPolygonWarehouse `json:"warehouse_location"`
}

type LinkDeliveryMethodToPolygonPolygon struct {
	// Polygon identifier
	PolygonId int64 `json:"polygon_id"`

	// Delivery time within polygon in minutes
	Time int64 `json:"time"`
}

type LinkDeliveryMethodToPolygonWarehouse struct {
	// Warehouse location latitude
	Latitude string `json:"lat"`

	// Warehouse location longitude
	Longitude string `json:"log"`
}

type LinkDeliveryMethodToPolygonResponse struct {
	core.CommonResponse
}

// Link delivery method to a delivery polygon
func (c Polygons) Link(params *LinkDeliveryMethodToPolygonParams) (*LinkDeliveryMethodToPolygonResponse, error) {
	url := "/v1/polygon/bind"

	resp := &LinkDeliveryMethodToPolygonResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DeletePolygonParams struct {
	// Polygons identifiers list
	PolygonIds []int64 `json:"polygon_ids"`
}

type DeletePolygonResponse struct {
	core.CommonResponse
}

// Delete polygon
func (c Polygons) Delete(params *DeletePolygonParams) (*DeletePolygonResponse, error) {
	url := "/v1/polygon/delete"

	resp := &DeletePolygonResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
