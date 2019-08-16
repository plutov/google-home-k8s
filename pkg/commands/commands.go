package commands

import (
	"errors"
	"fmt"

	"github.com/plutov/google-home-k8s/pkg/dialogflow"
	"github.com/plutov/google-home-k8s/pkg/k8s"
)

const (
	// ActionMain - when user starts a conversation
	ActionMain = "main"
	// ActionScaleReq - action to request statefulset/deployment scale
	ActionScaleReq = "scale_request"
	// ActionDoScale - action to do a scale after user provides new replicas count
	ActionDoScale = "do_scale"

	// MsgWelcome .
	MsgWelcome = "Welcome to Kubernetes Manager. How can I help you?"

	// ErrorScaleRequest .
	ErrorScaleRequest = "Sorry, I didn't get what exactly do you want to scale"
)

// Execute .
func Execute(k8sClient *k8s.Client, session *UserSession, req *dialogflow.Request) (string, error) {
	switch req.QueryResult.Action {
	case ActionMain:
		// Reset context
		session.ContextParams = nil

		return MsgWelcome, nil
	case ActionScaleReq:
		replicas, validateErr := validateScaleRequest(k8sClient, req)
		if validateErr != nil {
			return validateErr.Error(), nil
		}

		session.ContextParams = req.QueryResult.Parameters

		replicasText := formatReplicasNumber(replicas)
		return fmt.Sprintf("Got it. Currently, %s of the %s %s. To how many replicas do you want to scale?", replicasText, req.QueryResult.Parameters["resource_name"], req.QueryResult.Parameters["resource_type"]), nil
	}

	return "", errors.New("action is not supported")
}

func validateScaleRequest(k8sClient *k8s.Client, req *dialogflow.Request) (int32, error) {
	if _, ok := req.QueryResult.Parameters["resource_type"]; !ok || len(req.QueryResult.Parameters["resource_type"]) == 0 {
		return 0, errors.New(ErrorScaleRequest)
	}

	if _, ok := req.QueryResult.Parameters["resource_name"]; !ok || len(req.QueryResult.Parameters["resource_name"]) == 0 {
		return 0, errors.New(ErrorScaleRequest)
	}

	k8sResourceType, err := k8sClient.ValidateResourceType(req.QueryResult.Parameters["resource_type"])
	if err != nil {
		return 0, err
	}

	k8sResource, err := k8sClient.FindResourceByName(k8sResourceType, req.QueryResult.Parameters["resource_name"])
	if err != nil {
		return 0, err
	}

	req.QueryResult.Parameters["resource_type"] = k8sResourceType
	req.QueryResult.Parameters["resource_name"] = k8sResource.Name

	return k8sResource.Replicas, nil
}
