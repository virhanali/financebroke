package handler

import (
	"net/http"

	"financebroke/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	billUsecase usecase.BillUsecase
}

func NewDashboardHandler(billUsecase usecase.BillUsecase) *DashboardHandler {
	return &DashboardHandler{billUsecase: billUsecase}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userID := c.GetUint("user_id")

	dashboard, err := h.billUsecase.GetDashboard(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch dashboard"})
		return
	}

	c.JSON(http.StatusOK, dashboard)
}