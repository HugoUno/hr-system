package main

import (
	"hr-system/internal/handlers"
	"hr-system/internal/middleware"
	"hr-system/internal/services"
	"hr-system/pkg/cache"
	"hr-system/pkg/config"
	"hr-system/pkg/database"
	"hr-system/pkg/logger"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("無法加載配置: %v", err)
	}

	// 初始化日誌
	logger.InitLogger()

	// 初始化數據庫
	database.InitDB()

	// 初始化 Redis
	cache.InitRedis()

	// 初始化服務和處理器
	employeeService := services.NewEmployeeService()
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	leaveService := services.NewLeaveService()
	leaveHandler := handlers.NewLeaveHandler(leaveService)

	// 設置 Gin 路由
	r := gin.Default()

	// 添加中間件
	r.Use(middleware.ErrorHandler())
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 員工相關路由
		employees := api.Group("/employees")
		{
			employees.GET("/", employeeHandler.List)
			employees.POST("/", employeeHandler.Create)
			employees.GET("/:id", employeeHandler.Get)
			employees.PUT("/:id", employeeHandler.Update)
			employees.DELETE("/:id", employeeHandler.Delete)
		}

		// 請假相關路由
		leaves := api.Group("/leaves")
		{
			leaves.POST("/", leaveHandler.Create)
			leaves.GET("/employee/:employeeId", leaveHandler.List)
		}
	}

	// 啟動服務器
	r.Run(":" + config.Server.Port)
}
