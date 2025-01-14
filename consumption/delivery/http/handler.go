package http

import (
	"net/http"
	"re-home/auth/pkg/auth"
	consumption "re-home/consumption/usecase"
	"re-home/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase consumption.ConsumptionUseCase
}

func NewHandler(useCase consumption.ConsumptionUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	heating, err := h.useCase.GetConsumption(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"heating": heating})
}
