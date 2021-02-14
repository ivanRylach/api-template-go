package records

import (
    "context"
    "github.com/google/uuid"
    "go.uber.org/zap"
)

func (repo *RepositoryImpl) Write(ctx context.Context, record *RecordDTO) (created *RecordDTO, err error) {
    record.Id = uuid.NewString()

    result, err := repo.Collection.InsertOne(ctx, *record)

    if err == nil {
        zap.S().With("id", result.InsertedID).Info("Created a record")
        return record, nil
    } else {
        return nil, err
    }
}
