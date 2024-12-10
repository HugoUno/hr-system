# HR 系統後端服務

這是一個基於 Go 語言開發的人力資源管理系統後端服務，提供員工資訊管理和請假申請等功能。

## 系統需求

- Go 1.21 或以上
- Docker 和 Docker Compose
- Make

## 快速開始

1. 克隆專案
```bash
git clone https://github.com/HugoUno/hr-system.git
cd hr-system
```

2. 複製並修改環境配置文件
```bash
cp .env.example .env
# 根據需要修改 .env 中的配置
```

3. 啟動服務
```bash
# 構建並啟動所有服務
make deploy

# 或者分步驟執行：
make docker-build  # 構建 Docker 映像
make docker-up     # 啟動服務
make seed          # 填充測試數據
```

4. 驗證服務
```bash
# 測試 API 是否正常運行
curl http://localhost:8080/api/employees
```

## API 端點

### 員工管理
- `GET /api/employees` - 獲取所有員工列表
- `POST /api/employees` - 創建新員工
- `GET /api/employees/:id` - 獲取特定員工信息
- `PUT /api/employees/:id` - 更新員工信息
- `DELETE /api/employees/:id` - 刪除員工

### 請假管理
- `POST /api/leaves` - 提交請假申請
- `GET /api/leaves/employee/:employeeId` - 獲取員工的請假記錄

## 開發指南

### 專案結構
```
.
├── cmd/                    # 命令行工具
│   ├── migrate/           # 數據庫遷移工具
│   └── seed/              # 測試數據填充工具
├── internal/              # 內部程式碼
│   ├── models/           # 數據模型
│   ├── handlers/         # API 處理器
│   ├── middleware/       # 中間件
│   ├── repository/       # 數據訪問層
│   └── services/         # 業務邏輯層
├── pkg/                   # 公共包
│   ├── cache/            # Redis 緩存
│   ├── config/           # 配置管理
│   ├── database/         # 數據庫連接
│   └── logger/           # 日誌管理
├── scripts/              # SQL 腳本
├── tests/                # 測試文件
├── docker-compose.yml    # Docker 編排配置
├── Dockerfile           # Docker 構建文件
├── go.mod              # Go 模組文件
└── Makefile            # 構建和部署腳本
```

### 可用的 Make 命令
```bash
make build          # 構建應用程序
make run           # 本地運行應用程序
make test          # 運行單元測試
make docker-build  # 構建 Docker 映像
make docker-up     # 啟動所有 Docker 服務
make docker-down   # 停止所有 Docker 服務
make seed          # 填充測試數據
make deploy        # 完整部署流程
```

### 數據庫
- MySQL 用於持久化存儲
- Redis 用於緩存
- 初始化腳本位於 `scripts/init.sql`

### 配置說明
環境變量（在 .env 文件中配置）：
```env
DB_HOST=mysql
DB_USER=hruser
DB_PASSWORD=hrpassword
DB_NAME=hrdb
REDIS_HOST=redis
REDIS_PORT=6379
SERVER_PORT=8080
```

## 測試

### 運行測試
```bash
# 運行所有測試
make test

# 運行特定測試
go test ./tests/employee_test.go
```

## 常見問題

1. 服務無法啟動
   - 檢查 Docker 是否正在運行
   - 確認端口 8080, 3306, 6379 未被佔用
   - 檢查環境變量配置

2. 數據庫連接失敗
   - 確認 MySQL 服務已啟動
   - 檢查數據庫連接配置
   - 確認數據庫用戶權限

## 維護和支持

- 問題報告：請在 GitHub Issues 中提交
- 代碼貢獻：請提交 Pull Request