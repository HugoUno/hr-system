package tests

import (
	"hr-system/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// 使用測試數據庫配置
	dsn := "hruser:hrpassword@tcp(localhost:3306)/hrdb_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自動遷移測試數據庫結構
	err = db.AutoMigrate(&models.Employee{}, &models.LeaveRequest{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
