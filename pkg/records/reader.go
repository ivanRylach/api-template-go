package records

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *RepositoryImpl) Read(ctx context.Context, id string) (read *RecordDTO, err error) {
	cursor, err := repo.Collection.Find(ctx, bson.M{"id": id})

	if err == nil {
		var results []RecordDTO
		if err = cursor.All(ctx, &results); err == nil {
			if len(results) > 0 {
				return &results[0], nil
			}
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
