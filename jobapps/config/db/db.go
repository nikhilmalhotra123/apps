package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

// GetUserDB function
func GetUserDB() (*mongo.Collection, error) {
  ctx := context.TODO()

  client, err := mongo.NewClient( options.Client().ApplyURI("mongodb://mongo:27017"))
  if err != nil {
    return nil, err
  }

  err = client.Connect(ctx)
  if err != nil  {
    return nil, err
  }

  if err := client.Ping(ctx, nil); err != nil {
    return nil, err
  }

  collection := client.Database("GoLogin").Collection("users")
  return collection, nil

}

// GetApplicationDB function
func GetApplicationDB() (*mongo.Collection, error) {
  ctx := context.TODO()

  client, err := mongo.NewClient( options.Client().ApplyURI("mongodb://mongo:27017"))
  if err != nil {
    return nil, err
  }

  err = client.Connect(ctx)
  if err != nil  {
    return nil, err
  }

  if err := client.Ping(ctx, nil); err != nil {
    return nil, err
  }

  collection := client.Database("GoLogin").Collection("applications")
  return collection, nil

}
