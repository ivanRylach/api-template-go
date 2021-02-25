package records

import "encoding/json"

type RecordDTO struct {
	Id          string `json:"id" bson:"id"`
	Name        json.RawMessage `json:"name,string" bson:"name" binding:"required"`
	Description string `json:"description" bson:"description,omitempty" binding:"required"`
}
