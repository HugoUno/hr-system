package tests

import (
	"testing"

	"hr-system/internal/models"
	"hr-system/pkg/database"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// setupTestDB 初始化測試資料庫並回傳資料庫連線
func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	// 使用 SQLite 共享內存資料庫作為測試資料庫，無需額外服務
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	// 自動遷移測試資料庫結構
	if err := db.AutoMigrate(&models.Employee{}, &models.LeaveRequest{}); err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	// 將測試資料庫指派給全域變數，讓服務層能使用
	database.DB = db

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
		database.DB = nil
	})

	return db
}
