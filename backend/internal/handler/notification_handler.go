package handler

import (
	"fmt"
	"net/http"

	"financebroke/backend/internal/dto"
	"financebroke/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notificationUsecase: notificationUsecase}
}

func (h *NotificationHandler) UpdateNotificationSettings(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.NotificationSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.notificationUsecase.UpdateSettings(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update settings"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *NotificationHandler) TestTelegram(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.TestTelegramRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.notificationUsecase.TestTelegram(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send telegram message: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Telegram message sent successfully"})
}