package commands

import (
	"errors"
	"fmt"

	"github.com/plutov/google-home-k8s/pkg/dialogflow"
	"github.com/plutov/google-home-k8s/pkg/gke"
	log "github.com/sirupsen/logrus"
)

const (
	// ActionMain - when user starts a conversation
	ActionMain = "main"
	// ActionScaleReq - action to request statefulset/deployment scale
	ActionScaleReq = "scale_request"
	// ActionDoScale - action to do a scale after user provides new replicas count
	ActionDoScale = "do_scale"

	// ErrorClusterNoFound .
	ErrorClusterNoFound = "Kubernetes cluster not found, please check your Kubernetes settings."
)

// Execute .
func Execute(gkeClient *gke.Client, session *UserSession, req *dialogflow.Request) (string, error) {
	switch req.QueryResult.Action {
	case ActionMain:
		cluster, err := gkeClient.GetCluster()
		if err != nil {
			log.WithError(err).Error("unable to get cluster info")
			return ErrorClusterNoFound, nil
		}

		return fmt.Sprintf("Hi, you're currently in the %s Kubernetes cluster. How can I help you?", cluster.Name), nil
	case ActionScaleReq:
		session.ContextParams = "context"
		return "scale_request", nil
	}

	return "", errors.New("action is not supported")
}
