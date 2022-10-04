package controllers

import (
	"assignment_2/server/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, payload *views.Response) {
	ctx.JSON(payload.Status, payload)
}
