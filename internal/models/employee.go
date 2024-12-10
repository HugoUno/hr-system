package models

import (
	"time"

	"gorm.io/gorm"
)

// Employee 員工資訊模型
type Employee struct {
	gorm.Model
	EmployeeID   string    `gorm:"uniqueIndex;size:20"` // 員工編號
	Name         string    `gorm:"size:50"`             // 姓名
	Position     string    `gorm:"size:50"`             // 職位
	Department   string    `gorm:"size:50"`             // 部門
	Level        int       `gorm:"default:1"`           // 職等
	Salary       int       `gorm:"default:0"`           // 薪資
	Email        string    `gorm:"size:100"`            // 電子郵件
	Phone        string    `gorm:"size:20"`             // 聯絡電話
	Address      string    `gorm:"size:200"`            // 地址
	JoinDate     time.Time // 入職日期
	LeaveBalance int       `gorm:"default:0"` // 剩餘休假天數
	Status       string    `gorm:"size:20"`   // 在職狀態
}
