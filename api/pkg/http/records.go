package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ivanrylach.github.io/api/v1/pkg/records"
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
			return
		}

		created, err := (*r.Repo).Write(ctx, &payload)

		if err == nil {
			ctx.JSON(http.StatusCreated, *created)
			return
		} else {
			zap.S().Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorPayload{error: "Something went wrong in our side"})
			return
		}

	}
}

func (r *Records) Fetch() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var query records.RecordQuery
		if err := ctx.BindUri(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorPayload{error: "Malformed payload"})
			return
		}

		read, err := (*r.Repo).Read(ctx, query.Id)

		if err == nil {
			if read != nil {
				ctx.JSON(http.StatusOK, *read)
				return
			} else {
				ctx.JSON(http.StatusNotFound, "")
				return
			}
		} else {
			zap.S().Error(err)
			ctx.JSON(http.StatusInternalServerError, ErrorPayload{error: "Something went wrong in our side"})
			return
		}
	}
}
