# Dockerfile (开发环境)
FROM golang:1.25.4-trixie

# 安装 air 用于热重载
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 暴露端口
EXPOSE 8080

# 使用 air 运行（开发模式）
CMD ["air", "-c", ".air.toml"]