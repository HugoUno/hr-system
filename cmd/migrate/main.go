package main

import (
	"hr-system/internal/models"
	"hr-system/pkg/database"
	"log"
)

func main() {
	log.Println("開始數據庫遷移...")

	// 初始化數據庫連接
	database.InitDB()

	// 執行遷移
	err := database.DB.AutoMigrate(
		&models.Employee{},
		&models.LeaveRequest{},
	)
	if err != nil {
		log.Fatalf("遷移失敗: %v", err)
	}

	log.Println("數據庫遷移完成")
}
