package http

import (
	"net/http"
	consumption "re-home/consumption/usecase"
	"re-home/models"
	"strings"

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

func (h *Handler) GetConsumption(c *gin.Context) {
	//user := c.MustGet(auth.CtxUserKey).(string)

	filter := models.Filter(strings.ToLower(c.Query("filter")))

	// Valida l'azione
	// if !action.IsValid() {
	// 	http.Error(w, "Invalid action", http.StatusBadRequest)
	// 	return
	// }

	data, err := h.useCase.GetConsumption(c.Request.Context(), "", filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"consumption": data})
}
