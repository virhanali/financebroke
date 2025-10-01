package handler

import (
	"net/http"
	"strconv"

	"finance-app/backend/internal/dto"
	"finance-app/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type BillHandler struct {
	billUsecase usecase.BillUsecase
}

func NewBillHandler(billUsecase usecase.BillUsecase) *BillHandler {
	return &BillHandler{billUsecase: billUsecase}
}

func (h *BillHandler) CreateBill(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.BillCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bill, err := h.billUsecase.CreateBill(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill"})
		return
	}

	c.JSON(http.StatusCreated, bill)
}

func (h *BillHandler) GetBills(c *gin.Context) {
	userID := c.GetUint("user_id")

	bills, err := h.billUsecase.GetUserBills(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bills"})
		return
	}

	c.JSON(http.StatusOK, bills)
}

func (h *BillHandler) GetBill(c *gin.Context) {
	userID := c.GetUint("user_id")
	billID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	bill, err := h.billUsecase.GetBill(userID, uint(billID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}

	c.JSON(http.StatusOK, bill)
}

func (h *BillHandler) UpdateBill(c *gin.Context) {
	userID := c.GetUint("user_id")
	billID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	var req dto.BillUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bill, err := h.billUsecase.UpdateBill(userID, uint(billID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bill"})
		return
	}

	c.JSON(http.StatusOK, bill)
}

func (h *BillHandler) DeleteBill(c *gin.Context) {
	userID := c.GetUint("user_id")
	billID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	err = h.billUsecase.DeleteBill(userID, uint(billID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill deleted successfully"})
}

func (h *BillHandler) GetUpcomingBills(c *gin.Context) {
	userID := c.GetUint("user_id")

	bills, err := h.billUsecase.GetUpcomingBills(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch upcoming bills"})
		return
	}

	c.JSON(http.StatusOK, bills)
}