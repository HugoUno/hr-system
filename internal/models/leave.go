package models

import (
	"time"

	"gorm.io/gorm"
)

// LeaveType 請假類型
type LeaveType string

const (
	Annual    LeaveType = "annual"    // 年假
	Sick      LeaveType = "sick"      // 病假
	Personal  LeaveType = "personal"  // 事假
	Marriage  LeaveType = "marriage"  // 婚假
	Maternity LeaveType = "maternity" // 產假
)

// LeaveStatus 請假狀態
type LeaveStatus string

const (
	Pending   LeaveStatus = "pending"   // 待審核
	Approved  LeaveStatus = "approved"  // 已核准
	Rejected  LeaveStatus = "rejected"  // 已拒絕
	Cancelled LeaveStatus = "cancelled" // 已取消
)

// LeaveRequest 請假申請記錄
type LeaveRequest struct {
	gorm.Model
	EmployeeID   uint        `gorm:"index"` // 員工ID
	Employee     Employee    `gorm:"foreignKey:EmployeeID"`
	LeaveType    LeaveType   `gorm:"size:20"` // 請假類型
	StartDate    time.Time   // 開始日期
	EndDate      time.Time   // 結束日期
	Duration     float32     // 請假天數
	Reason       string      `gorm:"size:500"` // 請假原因
	Status       LeaveStatus `gorm:"size:20"`  // 請假狀態
	ApprovedBy   *uint       // 審核人ID
	ApprovedAt   *time.Time  // 審核時間
	ApproverNote string      `gorm:"size:500"` // 審核備註
}
