# HR 系統後端服務

這是一個簡單的人資系統後端服務，主要功能包含：
- 員工基本資料管理
- 請假申請與審核
- 薪資等級設定

## 開始使用

### 環境需求
- Go 1.21+
- Docker & Docker Compose
- Make

### 安裝步驟

1. 下載專案
```bash
git clone git@github.com:HugoUno/hr-system.git  # 建議用 SSH
cd hr-system
```

2. 設定環境變數
```bash
cp .env.example .env
vim .env  # 根據需求修改設定
```

3. 啟動服務
```bash
make deploy  # 這會自動執行所有必要步驟
```

4. 測試
```bash
curl localhost:8080/api/employees
# 應該會看到空陣列 [] 回傳
```

## API 文件

### 員工相關
- `GET /api/employees` - 取得員工列表
- `POST /api/employees` - 創建新員工
- `GET /api/employees/:id` - 獲取特定員工信息
- `PUT /api/employees/:id` - 更新員工信息
- `DELETE /api/employees/:id` - 刪除員工

### 請假相關
- `POST /api/leaves` - 提交請假申請
- `GET /api/leaves/employee/:employeeId` - 獲取員工的請假記錄

### 專案結構
```
.
├── cmd/                    # 主要命令
│   ├── migrate/           # DB 遷移
│   └── seed/              # 測試資料
├── internal/              # 私有程式碼
│   ├── models/           # 資料模型
│   ├── handlers/         # API 路由處理
│   ├── middleware/       # 中間件
│   ├── repository/       # 資料存取
│   └── services/         # 商業邏輯
├── pkg/                   # 共用程式碼
...
```

### Make 命令
```bash
- make build          # 構建應用程序
- make run           # 本地運行應用程序
+ make build         # 編譯程式
+ make run          # 本地執行
make test          # 測試
- make docker-build  # 構建 Docker 映像
+ make docker-build  # 建立容器
...
```