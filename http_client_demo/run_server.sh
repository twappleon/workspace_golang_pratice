#!/bin/bash

echo "=== HTTP模拟服务器 ==="
echo "正在启动服务器..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: 未找到Go命令，请先安装Go"
    exit 1
fi

# 进入server目录
cd "$(dirname "$0")/server"

# 安装依赖
echo "安装依赖..."
go mod tidy

# 运行服务器
echo "启动HTTP模拟服务器在端口8080..."
echo "按Ctrl+C停止服务器"
echo ""

go run . 