package records

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (repo *RepositoryImpl) Write(ctx context.Context, record *RecordDTO) (created *RecordDTO, err error) {
	record.Id = uuid.NewString()

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		result, err := repo.Collection.InsertOne(sessCtx, *record)
		if err != nil {
			return nil, err
		}

		return result.InsertedID, nil
	}

	session, err := repo.Client.Mongo.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	insertedId, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, err
	}

	zap.S().With("id", insertedId).Info("Created a record")
	return record, nil

}
