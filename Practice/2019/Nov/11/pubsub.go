package main

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	credentialJSONPath = "./secret/spolive-dev.json"
	projectID          = "spolive-dev"
	topicName          = "test"
)

func newClient(ctx context.Context) (*pubsub.Client, error) {

	jsonKey, err := ioutil.ReadFile(credentialJSONPath)
	conf, err := google.JWTConfigFromJSON(jsonKey, pubsub.ScopePubSub, pubsub.ScopeCloudPlatform)
	if err != nil {
		return nil, err
	}
	ts := conf.TokenSource(ctx)
	return pubsub.NewClient(ctx, projectID, option.WithTokenSource(ts))
}
