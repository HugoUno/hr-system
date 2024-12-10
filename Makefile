# 聲明所有的虛擬目標，確保這些命令能正確執行，即使有同名文件存在
# 包含了構建(build)、運行(run)、測試(test)、清理(clean)、Docker相關操作(docker-build/up/down)
# 以及數據庫遷移(migrate)和填充測試數據(seed)等任務
.PHONY: build run test clean docker-build docker-up docker-down migrate seed

# 構建應用程序
build:
	go build -o main .

# 運行應用程序
run:
	go run main.go

# 運行單元測試
test:
	go test -v ./...

# 清理構建文件
clean:
	rm -f main
	go clean

# 構建 Docker 映像檔
docker-build:
	docker-compose build

# 啟動所有 Docker 服務
docker-up:
	docker-compose up -d

# 停止所有 Docker 服務
docker-down:
	docker-compose down

# 填充測試數據
seed:
	go run cmd/seed/main.go

# 完整部署流程
deploy: docker-down docker-build docker-up seed 