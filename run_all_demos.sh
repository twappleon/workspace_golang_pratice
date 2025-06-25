#!/bin/bash

echo "=========================================="
echo "Go语言学习与实践工作空间 - 运行所有示例"
echo "=========================================="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 运行示例函数
run_demo() {
    local dir="$SCRIPT_DIR/$1"
    local name=$2
    local file=$3
    
    echo -e "\n${BLUE}运行 $name 示例...${NC}"
    echo "=========================================="
    
    if [ -d "$dir" ]; then
        cd "$dir"
        if go run "$file"; then
            echo -e "${GREEN}✓ $name 示例运行成功${NC}"
        else
            echo -e "${RED}✗ $name 示例运行失败${NC}"
        fi
        cd "$SCRIPT_DIR"
    else
        echo -e "${RED}✗ 目录 $dir 不存在${NC}"
    fi
}

# 运行基础概念示例
echo -e "\n${YELLOW}=== 基础概念示例 ===${NC}"
run_demo "init_demo" "包初始化" "init_demo.go"
run_demo "pass_by_value_demo" "值传递" "pass_by_value_demo.go"
run_demo "anonymous_var_demo" "匿名变量" "anonymous_var_demo.go"
run_demo "wait_done_demo" "Wait/Done模式" "wait_done_demo.go"
run_demo "channel_demo" "Channel基础" "channel_demo.go"
run_demo "sync_demo" "同步原语" "sync_demo.go"
run_demo "goroutine_leak_demo" "Goroutine泄漏" "goroutine_leak_demo.go"

# 运行并发编程示例
echo -e "\n${YELLOW}=== 并发编程示例 ===${NC}"
run_demo "concurrency_examples" "并发编程" "run_all_examples.sh"

echo -e "\n${GREEN}=========================================="
echo "所有示例运行完成！"
echo "==========================================${NC}"

echo -e "\n${YELLOW}提示：${NC}"
echo "1. 观察每个示例的输出，理解Go语言特性"
echo "2. 注意并发示例中结果的不确定性"
echo "3. 某些示例可能会产生goroutine泄漏（仅用于演示）"
echo "4. 使用 'go run -race' 可以检测竞态条件"
echo "5. 详细说明请查看各示例目录中的README.md文件" 