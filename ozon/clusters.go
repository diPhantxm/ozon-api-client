package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Clusters struct {
	client *core.Client
}

type ListClustersParams struct {
	// Clusters identifiers
	ClusterIds []string `json:"cluster_ids"`

	// Cluster type
	ClusterType string `json:"cluster_type"`
}

type ListClustersResponse struct {
	core.CommonResponse

	// Cluster details
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	// Cluster identifier
	Id int64 `json:"id"`

	// Cluster warehouse details
	LogisticClusters []LogisticCluster `json:"logistic_clusters"`

	// Cluster name
	Name string `json:"name"`

	// Cluster type
	Type string `json:"type"`
}

type LogisticCluster struct {
	// Warehouse status
	IsArchived bool `json:"is_archived"`

	// Warehouses
	Warehouses []LogisticClusterWarehouse `json:"warehouses"`
}

type LogisticClusterWarehouse struct {
	// Warehouse name
	Name string `json:"name"`

	// Warehouse type
	Type string `json:"type"`

	// Warehouse identifier
	Id int64 `json:"warehouse_id"`
}

func (c Clusters) List(ctx context.Context, params *ListClustersParams) (*ListClustersResponse, error) {
	url := "/v1/cluster/list"

	resp := &ListClustersResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
