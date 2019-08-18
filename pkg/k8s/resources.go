package k8s

import (
	"fmt"
	"strings"
)

const (
	// ResourceTypeDeployment .
	ResourceTypeDeployment = "deployment"
	// ResourceTypeStatefulSet .
	ResourceTypeStatefulSet = "stateful set"
	// ResourceTypeReplicaSet .
	ResourceTypeReplicaSet = "replica set"
)

var resourceTypes = map[string]string{
	"deployment":   ResourceTypeDeployment,
	"statefulset":  ResourceTypeStatefulSet,
	"stateful set": ResourceTypeStatefulSet,
	"replicaset":   ResourceTypeReplicaSet,
	"replica set":  ResourceTypeReplicaSet,
	"rs":           ResourceTypeReplicaSet,
}

// ValidateResourceType .
// Validates resource type
// Normalize it for Kubernetes
func (c *Client) ValidateResourceType(resourceType string) (string, error) {
	resourceType = strings.ToLower(resourceType)
	resourceType = strings.TrimSpace(resourceType)

	k8sName, ok := resourceTypes[resourceType]
	if !ok {
		return "", fmt.Errorf("%s is invalid resource type. Available options are: deployment, statefulset, replicaset", resourceType)
	}

	return k8sName, nil
}
