# --- 第一阶段：编译阶段 ---
FROM golang:1.25-alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

# 设置工作目录
WORKDIR /app

# 先复制 go.mod 和 go.sum 以利用缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码和生成的 swagger 文档
COPY . .

# 编译成可执行文件 (main.go 所在位置)
RUN go build -o main .

# --- 第二阶段：运行阶段 ---
FROM alpine:latest

# 安装基础库（处理时区等问题）
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /root/

# 从编译阶段复制二进制文件
COPY --from=builder /app/main .
# 复制配置文件和文档
COPY --from=builder /app/config.yaml ./
COPY --from=builder /app/docs ./docs

# 暴露端口（与 main.go 中一致）
EXPOSE 8080

# 启动程序
CMD ["./main"]