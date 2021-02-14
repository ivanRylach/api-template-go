package records

import "context"

type Repository interface {
    Write(ctx context.Context, record *RecordDTO) (*RecordDTO, error)
    Read(ctx context.Context, id string) (*RecordDTO, error)
}
