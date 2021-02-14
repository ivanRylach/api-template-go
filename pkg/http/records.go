package http

import (
	"api.ivanrylach.github.io/v1/pkg/records"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Records struct {
	Repo *records.Repository
}

func (r *Records) Create() func(*gin.Context) {
	return func(ctx *gin.Context) {
		var payload records.RecordDTO
		if err := ctx.BindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorPayload{error: "Malformed payload"})
		}

		created, err := (*r.Repo).Write(ctx, &payload)

		if err == nil {
			ctx.JSON(http.StatusCreated, *created)
		} else {
			zap.S().Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorPayload{error: "Something went wrong in our side"})
		}

	}
}

func (r *Records) Fetch() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var query records.RecordQuery
		if err := ctx.BindUri(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorPayload{error: "Malformed payload"})
		}

		read, err := (*r.Repo).Read(ctx, query.Id)

		if err == nil {
			if read != nil {
				ctx.JSON(http.StatusOK, *read)
			} else {
				ctx.JSON(http.StatusNotFound, "")
			}
		} else {
			zap.S().Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorPayload{error: "Something went wrong in our side"})
		}
	}
}
