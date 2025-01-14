package delivery

import (
	"re-home/auth/pkg/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, usecase auth.UseCase) {
	h := newHandler(usecase)

	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)
}
