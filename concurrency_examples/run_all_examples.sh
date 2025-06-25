#!/bin/bash

echo "=========================================="
echo "Go并发编程实践示例 - 运行所有示例"
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
run_example() {
    local dir="$SCRIPT_DIR/$1"
    local name=$2
    
    echo -e "\n${BLUE}运行 $name 示例...${NC}"
    echo "=========================================="
    
    if [ -d "$dir" ]; then
        cd "$dir"
        if go run main.go; then
            echo -e "${GREEN}✓ $name 示例运行成功${NC}"
        else
            echo -e "${RED}✗ $name 示例运行失败${NC}"
        fi
        cd "$SCRIPT_DIR"
    else
        echo -e "${RED}✗ 目录 $dir 不存在${NC}"
    fi
}

# 运行所有示例
run_example "race_condition" "竞态条件"
run_example "channel_behavior" "Channel行为"
run_example "memory_allocation" "内存分配"
run_example "goroutine_leak" "Goroutine泄漏"
run_example "context_demo" "Context使用"

echo -e "\n${GREEN}=========================================="
echo "所有示例运行完成！"
echo "==========================================${NC}"

echo -e "\n${YELLOW}提示：${NC}"
echo "1. 观察每个示例的输出，理解并发行为"
echo "2. 注意竞态条件示例中结果的不确定性"
echo "3. 某些示例可能会产生goroutine泄漏（仅用于演示）"
echo "4. 使用 'go run -race' 可以检测竞态条件" 