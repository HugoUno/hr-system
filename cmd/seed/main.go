package main

import (
	"hr-system/internal/models"
	"hr-system/pkg/database"
	"log"
	"time"
)

func main() {
	log.Println("開始填充測試數據...")

	// 初始化數據庫連接
	database.InitDB()

	// 創建測試員工數據
	employees := []models.Employee{
		{
			EmployeeID:   "EMP001",
			Name:         "張三",
			Position:     "軟體工程師",
			Department:   "研發部",
			Level:        3,
			Salary:       60000,
			Email:        "zhangsan@example.com",
			Phone:        "0912345678",
			Address:      "台北市信義區信義路100號",
			JoinDate:     time.Now().AddDate(-2, 0, 0),
			LeaveBalance: 14,
			Status:       "active",
		},
		{
			EmployeeID:   "EMP002",
			Name:         "李四",
			Position:     "產品經理",
			Department:   "產品部",
			Level:        4,
			Salary:       80000,
			Email:        "lisi@example.com",
			Phone:        "0923456789",
			Address:      "台北市大安區和平東路200號",
			JoinDate:     time.Now().AddDate(-1, -6, 0),
			LeaveBalance: 10,
			Status:       "active",
		},
	}

	// 批量創建員工記錄
	result := database.DB.Create(&employees)
	if result.Error != nil {
		log.Fatalf("創建員工數據失敗: %v", result.Error)
	}

	// 創建測試請假記錄
	leaveRequests := []models.LeaveRequest{
		{
			EmployeeID: 1,
			LeaveType:  models.Annual,
			StartDate:  time.Now().AddDate(0, 0, 5),
			EndDate:    time.Now().AddDate(0, 0, 7),
			Duration:   3,
			Reason:     "休假",
			Status:     models.Pending,
		},
		{
			EmployeeID: 2,
			LeaveType:  models.Sick,
			StartDate:  time.Now().AddDate(0, 0, 1),
			EndDate:    time.Now().AddDate(0, 0, 2),
			Duration:   2,
			Reason:     "感冒",
			Status:     models.Approved,
			ApprovedBy: &[]uint{1}[0],
			ApprovedAt: &[]time.Time{time.Now()}[0],
		},
	}

	// 批量創建請假記錄
	result = database.DB.Create(&leaveRequests)
	if result.Error != nil {
		log.Fatalf("創建請假記錄失敗: %v", result.Error)
	}

	log.Println("測試數據填充完成")
}
