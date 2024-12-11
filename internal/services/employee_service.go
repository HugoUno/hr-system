package services

import (
	"hr-system/internal/models"
	"hr-system/pkg/database"
)

type EmployeeService struct{}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (s *EmployeeService) List() ([]models.Employee, error) {
	var employees []models.Employee
	result := database.DB.Find(&employees)
	return employees, result.Error
}

func (s *EmployeeService) Create(employee *models.Employee) error {
	return database.DB.Create(employee).Error
}

func (s *EmployeeService) GetByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	result := database.DB.First(&employee, id)
	return &employee, result.Error
}

func (s *EmployeeService) Update(employee *models.Employee) error {
	return database.DB.Save(employee).Error
}

func (s *EmployeeService) Delete(id uint) error {
	return database.DB.Delete(&models.Employee{}, id).Error
}
