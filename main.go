package main

import (
	"os"

	"github.com/plutov/google-home-k8s/pkg/controllers"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}

func main() {
	handler, err := controllers.NewHandler()
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	r := controllers.NewRouter(handler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	if err := r.Start(":" + port); err != nil {
		r.Logger.Fatalf("shutting down the server: %s", err.Error())
	}
}
