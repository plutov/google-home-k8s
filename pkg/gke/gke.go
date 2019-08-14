package gke

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
)

// Client struct
type Client struct {
	svc       *container.Service
	projectID string
	zone      string
	clusterID string
}

var mandatoryEnvVars = []string{"GCP_PROJECT_ID", "GCP_ZONE", "GCP_CLUSTER_ID"}

// NewClient .
func NewClient() (*Client, error) {
	for _, envVarName := range mandatoryEnvVars {
		if len(os.Getenv(envVarName)) == 0 {
			return nil, fmt.Errorf("env var %s is required", envVarName)
		}
	}

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	svc, err := container.New(c)
	if err != nil {
		return nil, err
	}

	client := &Client{
		svc:       svc,
		projectID: os.Getenv("GCP_PROJECT_ID"),
		zone:      os.Getenv("GCP_ZONE"),
		clusterID: os.Getenv("GCP_CLUSTER_ID"),
	}

	return client, nil
}

// GetCluster .
func (k *Client) GetCluster() (*container.Cluster, error) {
	ctx := context.Background()
	resp, err := k.svc.Projects.Zones.Clusters.Get(k.projectID, k.zone, k.clusterID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
