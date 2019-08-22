[![Build Status](https://travis-ci.org/plutov/google-home-k8s.svg?branch=master)](https://travis-ci.org/plutov/google-home-k8s) [![GoDoc](https://godoc.org/github.com/plutov/google-home-k8s?status.svg)](https://godoc.org/github.com/plutov/google-home-k8s) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/google-home-k8s)](https://goreportcard.com/report/github.com/plutov/google-home-k8s)

# Google Home Action to communicate with Kubernetes

This project allows you to connect your Google Home to your Kubernetes cluster and control it via voice commands. It's not a public Google Home Action, since you have to configure Kubernetes access manually for your cluster.

This repository contains Dialogflow configuration as well, which can be imported into existing project. It can be customized later.

Example conversation:

> [you] Hey Google, talk to Kubernetes Manager

> [assistant] Welcome to Kubernetes Manager. How can I help you?

> [you] Scale statefulset "redis"

> [assistant] Got it. Currently, there are 3 replicas of the "redis" statefulset. To how many replicas do you want to scale?

> [you] 5

> [assistant] Statefulset has been updated. Anything else?

## Supported voice commands

> Scale statefulset "name"

> Scale deployment "name"

> Scale replicaset "name"

> What is the size of the cluster?

*Contribute to add more voice commands :)*

## Generate kubeconfig

To generate `kubeconfig` you have to install the following tools:
- `kubectl`
- `cfssl`
- `cfssljson`

Then run the following command to generate `build/kubeconfig` file.

```
./generate-kubeconfig.sh
```

## Deploy to Google App Engine

Requirements:
- `gcloud`

Set environment variables in `env.yaml` (`cp env.sample.yaml env.yaml`).

```
gcloud app deploy
```

Save URL, you will need to use it later in Dialogflow.

API deployed to App Engine is protected by static API Key which should be set in `env.yaml`. To access API, client should send `Authorization: Bearer ${API_KEY}` header.

## Configure in Dialogflow

1. Go to [Dialogflow Console](https://console.dialogflow.com/)
2. Select or create a new agent
3. Go to Settings -> Export and Import
4. Select **Import From Zip** (import this file [google-home-k8s.zip](https://raw.githubusercontent.com/plutov/google-home-k8s/master/google-home-k8s.zip))
5. Go to Fulfillment
6. Enable Webhook
7. Paste URL to API deployed to App Engine
8. Add Header. Key: `Authorization`, Value: `Bearer API_KEY` (replace `API_KEY` with the value from `env.yaml`)

## Unit Tests

```
GO111MODULE=on go test -mod vendor -race -v ./pkg/...
```