package main

import (
	"testing"
)

// TestBasicTypePassByValue 测试基本类型的值传递
func TestBasicTypePassByValue(t *testing.T) {
	original := 1
	modifyInt(original)

	// 原始值不应该改变（值传递）
	if original != 1 {
		t.Errorf("原始值应该保持不变，期望 1，实际 %d", original)
	}
}

// TestSlicePassByValue 测试切片的值传递
func TestSlicePassByValue(t *testing.T) {
	original := []int{1, 2, 3}
	modifySlice(original)

	// 切片是引用类型，原始值会改变
	if original[0] != 100 {
		t.Errorf("切片第一个元素应该被修改为 100，实际 %d", original[0])
	}

	// 长度应该保持不变（append 创建新描述符）
	if len(original) != 3 {
		t.Errorf("切片长度应该保持不变，期望 3，实际 %d", len(original))
	}
}

// TestPointerPassByValue 测试指针的值传递
func TestPointerPassByValue(t *testing.T) {
	original := 1
	modifyPointer(&original)

	// 通过指针修改，原始值应该改变
	if original != 100 {
		t.Errorf("通过指针修改后，原始值应该为 100，实际 %d", original)
	}
}

// TestNewPointer 测试 new 创建的指针
func TestNewPointer(t *testing.T) {
	p := new(int)
	*p = 1
	modifyPointer(p)

	if *p != 100 {
		t.Errorf("通过指针修改后，值应该为 100，实际 %d", *p)
	}
}

// TestSliceModification 测试切片修改的详细行为
func TestSliceModification(t *testing.T) {
	slice := []int{1, 2, 3}
	originalLen := len(slice)
	originalCap := cap(slice)

	modifySlice(slice)

	// 第一个元素应该被修改
	if slice[0] != 100 {
		t.Errorf("第一个元素应该为 100，实际 %d", slice[0])
	}

	// 其他元素应该保持不变
	if slice[1] != 2 || slice[2] != 3 {
		t.Errorf("其他元素应该保持不变")
	}

	// 长度和容量应该保持不变（append 创建新描述符）
	if len(slice) != originalLen {
		t.Errorf("长度应该保持不变，期望 %d，实际 %d", originalLen, len(slice))
	}

	if cap(slice) != originalCap {
		t.Errorf("容量应该保持不变，期望 %d，实际 %d", originalCap, cap(slice))
	}
}

// BenchmarkBasicTypePass 基准测试基本类型传递
func BenchmarkBasicTypePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		modifyInt(i)
	}
}

// BenchmarkSlicePass 基准测试切片传递
func BenchmarkSlicePass(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		modifySlice(slice)
	}
}

// BenchmarkPointerPass 基准测试指针传递
func BenchmarkPointerPass(b *testing.B) {
	value := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		modifyPointer(&value)
	}
}
