package mongodb

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.uber.org/zap"
    "time"
)

type Client struct {
    Mongo *mongo.Client
}

func NewClient() *Client {
    creds := options.Credential{
        Username: "root",
        Password: "example",
    }
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(creds))
    if err != nil {
        zap.S().Panic(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        zap.S().Panic(err)
    }

    return &Client{Mongo: client}
}

func (c *Client) Shutdown(ctx context.Context) {
    if err := c.Mongo.Disconnect(ctx); err != nil {
        zap.S().Panic(err)
    }
}
