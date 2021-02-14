package records

type RecordQuery struct {
    Id string `uri:"id" binding:"required,uuid"`
}
