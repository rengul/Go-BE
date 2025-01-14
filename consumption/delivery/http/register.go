package http

import (
	"re-home/consumption/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc usecase.ConsumptionUseCase) {
	h := NewHandler(uc)

	consumption := router.Group("/consumption")
	{
		consumption.GET("", h.Get)
	}
}
