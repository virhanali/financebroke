package handler

import (
	"net/http"
	"strconv"

	"financebroke/backend/internal/dto"
	"financebroke/backend/internal/usecase"
	"financebroke/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	logger := utils.GetLogger()
	logger.Info("[HANDLER] Register attempt")

	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogError("AUTH_HANDLER_BIND", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid request format",
			"details":    err.Error(),
			"error_code": "INVALID_REQUEST",
		})
		return
	}

	logger.Info("[HANDLER] Register request validated", map[string]interface{}{
		"email": req.Email,
		"name":  req.Name,
	})

	response, err := h.authUsecase.Register(&req)
	if err != nil {
		utils.LogError("AUTH_HANDLER_REGISTER", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Failed to create user",
			"details":    err.Error(),
			"error_code": "USER_CREATION_FAILED",
		})
		return
	}

	logger.Info("[HANDLER] User created successfully", map[string]interface{}{
		"user_id": response.User.ID,
		"email":   response.User.Email,
	})
	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	response, err := h.authUsecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid credentials",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userIDStr := c.GetUint("user_id")
	userID, err := strconv.ParseUint(strconv.FormatUint(uint64(userIDStr), 10), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.authUsecase.GetProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
