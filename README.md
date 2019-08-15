[![Build Status](https://travis-ci.org/plutov/google-home-k8s.svg?branch=master)](https://travis-ci.org/plutov/google-home-k8s) [![GoDoc](https://godoc.org/github.com/plutov/google-home-k8s?status.svg)](https://godoc.org/github.com/plutov/google-home-k8s) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/google-home-k8s)](https://goreportcard.com/report/github.com/plutov/google-home-k8s)

# Google Home Action to communicate with Kubernetes

This project allows you to connect your Google Home to your Kubernetes (only GKE for now) cluster and control it via voice commands. It's not a public Google Home Action, since you have to configure GKE access manually.

This repository contains Dialogflow configuration as well, which can be imported into existing project. It can be customized later.

Example conversation:

> [you] Hey Google, talk to Kubernetes Manager

> [assistant] Hi, you're currently in the "sandbox" Kubernetes cluster. How can I help you?

> [you] Scale statefulset "redis"

> [assistant] Got it. Currently, there are 3 replicas of the "redis" statefulset. To how many replicas do you want to scale?

> [you] 5

> [assistant] Done. Anything else?

## Supported voice commands

> Scale statefulset "name"

> Scale deployment "name"

*Contribute to add more voice commands :)*

## Deploy to Google App Engine

Requirements:
- `gcloud`

Set environment variables in `env.yaml` (`cp env.sample.yaml env.yaml`).

```
gcloud app deploy
```

Save API URL, you will need to use it later in Dialogflow.

API deployed to App Engine is protected by static API Key which should be set in `env.yaml`. To access API, client should send `Authorization: Bearer ${API_KEY}` header.

## Give App Engine access to GKE

1. Go to IAM
2. Find GAE service account, which ends with `@appspot.gserviceaccount.com`
3. Grant `Kubernetes Engine Developer` role

## Configure in Dialogflow

1. Go to [Dialogflow Console](https://console.dialogflow.com/)
2. Select or create a new agent
3. Go to Settings -> Export and Import
4. Select **Import From Zip** (import this file [dialogflow.zip](https://raw.githubusercontent.com/plutov/google-home-k8s/master/dialogflow.zip))
5. Go to Fulfillment
6. Enable Webhook
7. Paste URL to API deployed to App Engine
8. Add Header. Key: `Authorization`, Value: `Bearer API_KEY` (replace `API_KEY` with the value from `env.yaml`)
