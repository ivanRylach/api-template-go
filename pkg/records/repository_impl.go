package records

import (
	"api.ivanrylach.github.io/v1/pkg/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.uber.org/zap"
	"time"
)

type RepositoryImpl struct {
	Client     *mongodb.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongodb.Client) *RepositoryImpl {
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	collection := client.Mongo.Database("api-template").Collection("records", wcMajorityCollectionOpts)

	createIndex(context.Background(), collection, "id", true)

	return &RepositoryImpl{
		Client:     client,
		Collection: collection,
	}
}

func createIndex(ctx context.Context, collection *mongo.Collection, field string, unique bool) {

	indexModel := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
		zap.S().Panic(err)
	}
}
