package records

type RecordDTO struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name" binding:"required"`
	Description string `json:"description" bson:"description,omitempty" binding:"required"`
}
