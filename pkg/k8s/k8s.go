package k8s

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Client struct
type Client struct {
	apiClient *kubernetes.Clientset
}

// CommonResourceInfo - we use this type to combine specs of different resources
type CommonResourceInfo struct {
	Name     string
	Replicas int32
}

// NewClient .
func NewClient() (*Client, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "./build/kubeconfig")
	if err != nil {
		return nil, err
	}

	apiClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		apiClient: apiClient,
	}

	return client, nil
}

// FindResourceByName .
// Finds resource by name
// Returns k8s resource specs (name could be different from the one specified by the user)
func (c *Client) FindResourceByName(resourceType string, resourceName string) (*CommonResourceInfo, error) {
	userName := normalizeName(resourceName)
	notFoundErr := fmt.Errorf("%s %s not found", resourceType, resourceName)

	var (
		list []CommonResourceInfo
		err  error
	)

	switch resourceType {
	case ResourceTypeDeployment:
		list, err = c.getDeploymentList()
		break
	case ResourceTypeStatefulSet:
		list, err = c.getStatefulsetList()
		break
	case ResourceTypeReplicaSet:
		list, err = c.getRSList()
		break
	}

	if err != nil {
		log.WithError(err).Error("unable to get list of resources")
		return nil, notFoundErr
	}

	for _, i := range list {
		k8sName := normalizeName(i.Name)
		if k8sName == userName {
			return &i, nil
		}
	}

	return nil, notFoundErr
}

// ScaleResource .
func (c *Client) ScaleResource(resourceType string, resourceName string, replicas int32) error {
	ctx := context.Background()

	notFoundErr := fmt.Errorf("%s %s not found", resourceType, resourceName)

	switch resourceType {
	case ResourceTypeDeployment:
		client := c.apiClient.AppsV1().Deployments(os.Getenv("NAMESPACE"))
		res, err := client.Get(ctx, resourceName, metav1.GetOptions{})
		if err != nil {
			return notFoundErr
		}

		res.Spec.Replicas = &replicas
		_, updateErr := client.Update(ctx, res, metav1.UpdateOptions{})
		return updateErr
	case ResourceTypeStatefulSet:
		client := c.apiClient.AppsV1().StatefulSets(os.Getenv("NAMESPACE"))
		res, err := client.Get(ctx, resourceName, metav1.GetOptions{})
		if err != nil {
			return notFoundErr
		}

		res.Spec.Replicas = &replicas
		_, updateErr := client.Update(ctx, res, metav1.UpdateOptions{})
		return updateErr
	case ResourceTypeReplicaSet:
		client := c.apiClient.AppsV1().ReplicaSets(os.Getenv("NAMESPACE"))
		res, err := client.Get(ctx, resourceName, metav1.GetOptions{})
		if err != nil {
			return notFoundErr
		}

		res.Spec.Replicas = &replicas
		_, updateErr := client.Update(ctx, res, metav1.UpdateOptions{})
		return updateErr
	}

	return nil
}

// GetClusterSize .
func (c *Client) GetClusterSize() (int, error) {
	ctx := context.Background()

	nodes, err := c.apiClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	return len(nodes.Items), err
}

func (c *Client) getDeploymentList() ([]CommonResourceInfo, error) {
	ctx := context.Background()

	client := c.apiClient.AppsV1().Deployments(os.Getenv("NAMESPACE"))
	res, err := client.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var list []CommonResourceInfo
	for _, i := range res.Items {
		list = append(list, CommonResourceInfo{
			Name:     i.Name,
			Replicas: *i.Spec.Replicas,
		})
	}

	return list, nil
}

func (c *Client) getStatefulsetList() ([]CommonResourceInfo, error) {
	ctx := context.Background()

	client := c.apiClient.AppsV1().StatefulSets(os.Getenv("NAMESPACE"))
	res, err := client.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var list []CommonResourceInfo
	for _, i := range res.Items {
		list = append(list, CommonResourceInfo{
			Name:     i.Name,
			Replicas: *i.Spec.Replicas,
		})
	}

	return list, nil
}

func (c *Client) getRSList() ([]CommonResourceInfo, error) {
	ctx := context.Background()

	client := c.apiClient.AppsV1().ReplicaSets(os.Getenv("NAMESPACE"))
	res, err := client.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var list []CommonResourceInfo
	for _, i := range res.Items {
		list = append(list, CommonResourceInfo{
			Name:     i.Name,
			Replicas: *i.Spec.Replicas,
		})
	}

	return list, nil
}
