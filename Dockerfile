# 使用 Go 1.21 Alpine 作為基礎鏡像
FROM golang:1.21-alpine

# 設置工作目錄
WORKDIR /app

# 安裝必要的系統依賴
RUN apk add --no-cache make gcc musl-dev

# 複製並下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製源代碼
COPY . .

# 構建應用
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 運行應用
CMD ["./main"]
