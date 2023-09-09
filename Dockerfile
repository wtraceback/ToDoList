# 使用基于 golang 的 Docker 镜像
FROM golang:latest

# 构建者的基本信息
MAINTAINER whxcer

# 为镜像设置必要的环境变量(避免在构建镜像的时候拉取开源库超时)
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 将本地代码复制到容器中的工作目录
COPY . .

# 构建 Go 应用
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 运行 Go 应用
CMD ["./main"]