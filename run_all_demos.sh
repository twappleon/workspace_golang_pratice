#!/bin/bash

# Go 语言实践项目 - 运行所有演示脚本
# 这个脚本会依次运行所有 Go 演示程序

echo "🚀 开始运行所有 Go 语言演示程序..."
echo "=================================="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 运行单个演示的函数
run_demo() {
    local demo_name=$1
    local demo_dir=$2
    local go_file=$3
    
    echo -e "\n${BLUE}📁 运行 $demo_name...${NC}"
    echo "----------------------------------------"
    
    if [ -d "$demo_dir" ] && [ -f "$demo_dir/$go_file" ]; then
        cd "$demo_dir"
        if go run "$go_file"; then
            echo -e "${GREEN}✅ $demo_name 运行成功${NC}"
        else
            echo -e "${RED}❌ $demo_name 运行失败${NC}"
        fi
        cd ..
    else
        echo -e "${YELLOW}⚠️  $demo_name 目录或文件不存在${NC}"
    fi
}

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go 未安装，请先安装 Go${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go 版本: $(go version)${NC}"

# 运行所有演示
run_demo "Init 函数演示" "init_demo" "init_demo.go"
run_demo "Goroutine 泄漏演示" "goroutine_leak_demo" "goroutine_leak_demo.go"
run_demo "值传递演示" "pass_by_value_demo" "pass_by_value_demo.go"
run_demo "匿名变量演示" "anonymous_var_demo" "anonymous_var_demo.go"
run_demo "WaitGroup 演示" "wait_done_demo" "wait_done_demo.go"
run_demo "Channel 演示" "channel_demo" "channel_demo.go"
run_demo "同步原语演示" "sync_demo" "sync_demo.go"
run_demo "压测演示" "benchmark_demo" "benchmark_demo.go"

echo -e "\n${GREEN}🎉 所有演示程序运行完成！${NC}"
echo "=================================="
echo -e "${BLUE}💡 提示：${NC}"
echo "   - 使用 'go run <文件名>' 运行单个程序"
echo "   - 使用 'go build' 构建可执行文件"
echo "   - 使用 'go test' 运行测试"
echo "   - 查看各目录下的 README.md 了解详细信息"
echo "   - 压测工具位于 benchmark_tools/ 目录" 