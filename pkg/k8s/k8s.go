package k8s

import (
	"fmt"
	"os"

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

func (c *Client) getDeploymentList() ([]CommonResourceInfo, error) {
	client := c.apiClient.AppsV1().Deployments(os.Getenv("NAMESPACE"))
	res, err := client.List(metav1.ListOptions{})
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
	client := c.apiClient.AppsV1().StatefulSets(os.Getenv("NAMESPACE"))
	res, err := client.List(metav1.ListOptions{})
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
	client := c.apiClient.AppsV1().ReplicaSets(os.Getenv("NAMESPACE"))
	res, err := client.List(metav1.ListOptions{})
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
