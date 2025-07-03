#!/bin/bash

echo "=== HTTP客户端调用示例 ==="
echo "正在启动演示程序..."

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

# 运行演示
echo "运行HTTP客户端演示..."
echo "程序将自动启动本地服务器并执行各种HTTP请求示例"
echo "按Ctrl+C停止程序"
echo ""

go run . 