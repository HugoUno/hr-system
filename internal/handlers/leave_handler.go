package handlers

import (
	"net/http"
	"strconv"

	"hr-system/internal/models"
	"hr-system/internal/services"

	"github.com/gin-gonic/gin"
)

type LeaveHandler struct {
	service *services.LeaveService
}

func NewLeaveHandler(service *services.LeaveService) *LeaveHandler {
	return &LeaveHandler{service: service}
}

// 創建請假申請
func (h *LeaveHandler) Create(c *gin.Context) {
	var leave models.LeaveRequest
	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateLeaveRequest(&leave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, leave)
}

// 獲取員工的請假記錄
func (h *LeaveHandler) List(c *gin.Context) {
	employeeID, err := strconv.ParseUint(c.Param("employeeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	leaves, err := h.service.GetLeaveRequests(uint(employeeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaves)
}
