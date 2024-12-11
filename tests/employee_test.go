package tests

import (
	"testing"

	"hr-system/internal/models"
	"hr-system/internal/services"
)

func TestCreateEmployee(t *testing.T) {
	// TODO: 實現員工創建測試
}

func TestLeaveRequest(t *testing.T) {
	// TODO: 實現請假申請測試
}

func TestEmployeeService_Create(t *testing.T) {
	// 設置測試數據庫
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("設置測試數據庫失敗: %v", err)
	}

	service := services.NewEmployeeService()
	employee := &models.Employee{
		Name:     "測試員工",
		Position: "工程師",
	}

	err = service.Create(employee)
	if err != nil {
		t.Errorf("創建員工失敗: %v", err)
	}

	// 驗證結果
	var result models.Employee
	err = db.First(&result, employee.ID).Error
	if err != nil {
		t.Errorf("查詢員工失敗: %v", err)
	}

	if result.Name != employee.Name {
		t.Errorf("員工名稱不匹配，期望 %s，得到 %s", employee.Name, result.Name)
	}
}
