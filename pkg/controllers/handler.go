package controllers

import (
	"github.com/plutov/google-home-k8s/pkg/commands"
	"github.com/plutov/google-home-k8s/pkg/k8s"
)

// Handler type
type Handler struct {
	KubernetesClient   *k8s.Client
	UserSessionManager *commands.UserSessionManager
}

// NewHandler constructor
func NewHandler() (*Handler, error) {
	h := new(Handler)

	var err error
	h.KubernetesClient, err = k8s.NewClient()
	if err != nil {
		return nil, err
	}

	h.UserSessionManager = commands.NewUserSessionManager()

	return h, nil
}
