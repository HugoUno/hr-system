package services

import (
	"hr-system/internal/models"
	"hr-system/pkg/database"
)

type LeaveService struct{}

func NewLeaveService() *LeaveService {
	return &LeaveService{}
}

func (s *LeaveService) CreateLeaveRequest(leave *models.LeaveRequest) error {
	return database.DB.Create(leave).Error
}

func (s *LeaveService) GetLeaveRequests(employeeID uint) ([]models.LeaveRequest, error) {
	var leaves []models.LeaveRequest
	result := database.DB.Where("employee_id = ?", employeeID).Find(&leaves)
	return leaves, result.Error
}

func (s *LeaveService) UpdateLeaveStatus(id uint, status models.LeaveStatus, approverID uint) error {
	return database.DB.Model(&models.LeaveRequest{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      status,
			"approved_by": approverID,
		}).Error
}
