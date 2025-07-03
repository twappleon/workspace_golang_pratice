#!/bin/bash

echo "=== HTTP客户端演示 ==="
echo "正在启动客户端..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: 未找到Go命令，请先安装Go"
    exit 1
fi

# 进入项目目录
cd "$(dirname "$0")"

# 安装依赖
echo "安装依赖..."
go mod tidy

# 运行客户端演示
echo "运行HTTP客户端演示..."
echo "请确保服务器已在端口8080运行"
echo ""

go run . 