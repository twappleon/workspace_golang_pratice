#!/bin/bash

# 运行所有测试的脚本
echo "开始运行所有测试..."

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 测试结果统计
total_tests=0
passed_tests=0
failed_tests=0

# 运行测试的函数
run_test() {
    local dir=$1
    local test_name=$2
    
    echo -e "\n${YELLOW}运行测试: $test_name${NC}"
    echo "目录: $dir"
    
    cd "$dir" || {
        echo -e "${RED}错误: 无法进入目录 $dir${NC}"
        return 1
    }
    
    # 运行测试
    if go test -v ./...; then
        echo -e "${GREEN}✓ $test_name 测试通过${NC}"
        ((passed_tests++))
    else
        echo -e "${RED}✗ $test_name 测试失败${NC}"
        ((failed_tests++))
    fi
    
    ((total_tests++))
    cd - > /dev/null
}

# 运行基准测试的函数
run_benchmark() {
    local dir=$1
    local benchmark_name=$2
    
    echo -e "\n${YELLOW}运行基准测试: $benchmark_name${NC}"
    echo "目录: $dir"
    
    cd "$dir" || {
        echo -e "${RED}错误: 无法进入目录 $dir${NC}"
        return 1
    }
    
    # 运行基准测试
    if go test -bench=. -benchmem ./...; then
        echo -e "${GREEN}✓ $benchmark_name 基准测试完成${NC}"
    else
        echo -e "${RED}✗ $benchmark_name 基准测试失败${NC}"
    fi
    
    cd - > /dev/null
}

# 检查是否有测试文件
check_test_files() {
    local dir=$1
    if [ -d "$dir" ] && [ -f "$dir"/*_test.go ]; then
        return 0
    fi
    return 1
}

# 主测试流程
echo "检查测试文件..."

# 运行各个模块的测试
if check_test_files "init_demo"; then
    run_test "init_demo" "Init 函数测试"
fi

if check_test_files "channel_demo"; then
    run_test "channel_demo" "Channel 测试"
fi

if check_test_files "sync_demo"; then
    run_test "sync_demo" "同步原语测试"
fi

if check_test_files "pass_by_value_demo"; then
    run_test "pass_by_value_demo" "值传递测试"
fi

if check_test_files "wait_done_demo"; then
    run_test "wait_done_demo" "WaitGroup 测试"
fi

if check_test_files "anonymous_var_demo"; then
    run_test "anonymous_var_demo" "匿名变量测试"
fi

if check_test_files "benchmark_demo"; then
    run_test "benchmark_demo" "压测工具测试"
fi

if check_test_files "channel_no_close_demo"; then
    run_test "channel_no_close_demo" "Channel 关闭测试"
fi

# 运行 concurrency_examples 中的测试
if [ -d "concurrency_examples" ]; then
    echo -e "\n${YELLOW}运行并发示例测试...${NC}"
    
    for subdir in concurrency_examples/*/; do
        if [ -d "$subdir" ] && check_test_files "$subdir"; then
            test_name=$(basename "$subdir")
            run_test "$subdir" "并发示例 - $test_name"
        fi
    done
fi

# 运行基准测试（可选）
echo -e "\n${YELLOW}是否运行基准测试? (y/n)${NC}"
read -r run_benchmarks

if [[ $run_benchmarks =~ ^[Yy]$ ]]; then
    echo -e "\n${YELLOW}开始运行基准测试...${NC}"
    
    if check_test_files "channel_demo"; then
        run_benchmark "channel_demo" "Channel 基准测试"
    fi
    
    if check_test_files "sync_demo"; then
        run_benchmark "sync_demo" "同步原语基准测试"
    fi
    
    if check_test_files "benchmark_demo"; then
        run_benchmark "benchmark_demo" "压测工具基准测试"
    fi
fi

# 输出测试结果统计
echo -e "\n${YELLOW}=== 测试结果统计 ===${NC}"
echo -e "总测试数: $total_tests"
echo -e "${GREEN}通过: $passed_tests${NC}"
echo -e "${RED}失败: $failed_tests${NC}"

if [ $failed_tests -eq 0 ]; then
    echo -e "\n${GREEN}🎉 所有测试通过！${NC}"
    exit 0
else
    echo -e "\n${RED}❌ 有 $failed_tests 个测试失败${NC}"
    exit 1
fi 