package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// DEFAULT_CAPACITY 默认容量
const DEFAULT_CAPACITY = 10

// CustomizedArray 动态数组结构体参数
type CustomizedArray struct {
	// 数组
	array []int
	// 数组大小
	size int
}

// NewCustomizedArray 初始化动态数组
func NewCustomizedArray() *CustomizedArray {
	// 初始化参数
	return &CustomizedArray{
		// 原来的数组不需要删除，Go 中有 GC 垃圾自动回收机制
		array: make([]int, DEFAULT_CAPACITY),
		size:  0,
	}
}

// 添加元素
func (arr *CustomizedArray) add(value int) {
	if arr.size == DEFAULT_CAPACITY {
		// 扩容机制
		arr.expandCapacity()
	}
	arr.array[arr.size] = value
	arr.size++
}

// 获取元素
func (arr *CustomizedArray) get(index int) int {
	if arr.size == 0 || index < 0 || index > arr.size {
		return math.MinInt
	}
	return arr.array[index]
}

// 更新元素
func (arr *CustomizedArray) update(index int, value int) bool {
	if arr.size == 0 || index < 0 || index > arr.size {
		return false
	}
	arr.array[index] = value
	return true
}

// 移除元素
func (arr *CustomizedArray) remove(index int) int {
	if arr.size == 0 || index < 0 || index > arr.size {
		return math.MinInt
	}
	// 记录这个结果
	result := arr.array[index]
	for i := index + 1; i < arr.size; i++ {
		arr.array[i] = arr.array[i+1]
	}
	arr.size--
	return result
}

// Size 获取数组大小
func (arr *CustomizedArray) Size() int {
	return arr.size
}

// 扩容操作
func (arr *CustomizedArray) expandCapacity() {
	newArr := make([]int, len(arr.array)*2)
	// 这个地方还是使用了库函数
	copy(newArr, arr.array)
	arr.array = newArr
}

// 手写扩容机制
func (arr *CustomizedArray) expandCapacityHand() {
	// 创建一个新的数组，容量是原来的两倍
	newArr := make([]int, len(arr.array)*2)

	// 将愿数组中的元素逐个复制到新数组中
	for i := 0; i < len(arr.array); i++ {
		newArr[i] = arr.array[i]
	}

	// 将愿数组中的引用指向新数组，不需要将愿数组删除，Go GC垃圾回收机制会去进行处理
	arr.array = newArr
}

// ToString 打印数组信息
func (arr *CustomizedArray) ToString() string {
	var result strings.Builder
	result.WriteString("数组里面的元素为：")

	for i := 0; i < arr.size; i++ {
		result.WriteString(strconv.Itoa(arr.array[i]))
		result.WriteString(" ")
	}

	return result.String()
}

func main() {
	myArr := &CustomizedArray{
		array: make([]int, DEFAULT_CAPACITY),
		size:  0,
	}
	myArr.add(1)
	myArr.add(2)
	fmt.Println(myArr.ToString())
	fmt.Println(myArr.get(0))
	fmt.Println(myArr.Size())
	myArr.update(0, -1)
	fmt.Println(myArr.ToString())
	myArr.remove(1)
	fmt.Println(myArr.ToString())
}
