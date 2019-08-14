package controllers

import (
	"github.com/plutov/google-home-k8s/pkg/gke"
)

// Handler type
type Handler struct {
	GKEClient *gke.Client
}

// NewHandler constructor
func NewHandler() (*Handler, error) {
	h := new(Handler)

	var err error
	h.GKEClient, err = gke.NewClient()
	if err != nil {
		return nil, err
	}

	return h, nil
}
