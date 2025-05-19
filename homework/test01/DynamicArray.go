package main

import (
	"math"
	"strconv"
	"strings"
)

/**
6.1 API 的设计 API (Application Programming Interface - 程序编程接口)
 *  1. 添加元素 add
 *  2. 访问元素 get
 *  3. 修改元素 update
 *  4. 删除元素 remove
 *  5. 大小 size
*/

// DEFAULT_CAPACITY 默认数组容量
const DEFAULT_CAPACITY = 10

// DynamicArray 创建动态数组结构体
type DynamicArray struct {
	// 数组结构
	array []int
	// 大小
	size int
}

// NewDynamicArray 初始化动态数组
func NewDynamicArray() *DynamicArray {
	// 进行初始化
	return &DynamicArray{
		array: make([]int, DEFAULT_CAPACITY),
		size:  0,
	}
}

// add 添加元素 API
func (arr *DynamicArray) add(value int) {
	// 达到阈值
	if arr.size == DEFAULT_CAPACITY {
		// 扩容机制
		//arr.expandArray()
	}
	arr.array[arr.size] = value
	arr.size++
}

// get 访问元素 API
func (arr *DynamicArray) get(index int) int {
	// 特判
	if arr.size == 0 || index < 0 || index >= arr.size {
		// 设置最小值，予以区分
		return math.MinInt
	}
	return arr.array[index]
}

// update 更新元素 API
func (arr *DynamicArray) update(index int, value int) bool {
	// 特判
	if arr.size == 0 || index < 0 || index >= arr.size {
		return false
	}
	arr.array[index] = value
	return true
}

// remove 移除元素 API
func (arr *DynamicArray) remove(index int) int {
	// 特判
	if arr.size == 0 || index < 0 || index >= arr.size {
		return math.MinInt
	}
	// 记录这个结果
	result := arr.array[index]
	// 查找到删除的元素，其他元素进行位移
	for i := index + 1; i < arr.size-1; i++ {
		arr.array[i] = arr.array[i+1]
	}
	arr.size--
	return result
}

// Size 获取数组大小
func (arr *DynamicArray) Size() int {
	return arr.size
}

// expandArray 手写扩容函数
func (arr *DynamicArray) expandArray() {
	// 初始化长度是原本数组两倍的新数组
	newArr := make([]int, len(arr.array)*2)

	// 将愿数组中的元素逐个复制到新数组中
	for i := 0; i < len(arr.array); i++ {
		newArr[i] = arr.array[i]
	}
	// 将愿数组中的引用指向新数组，不需要将愿数组删除，Go GC垃圾回收机制会去进行处理
	arr.array = newArr
}

// ToString 打印数组信息
func (arr *DynamicArray) ToString() string {
	var result strings.Builder
	result.WriteString("数组里面的元素为：")

	for i := 0; i < arr.size; i++ {
		result.WriteString(strconv.Itoa(arr.array[i]))
		result.WriteString(" ")
	}

	return result.String()
}

//func main() {
//	myArr := NewDynamicArray()
//	myArr.add(1)
//	myArr.add(2)
//	fmt.Println(myArr.ToString())
//	fmt.Println(myArr.get(0))
//	fmt.Println(myArr.Size())
//	myArr.update(0, -1)
//	fmt.Println(myArr.ToString())
//	removed := myArr.remove(1)
//	fmt.Println(myArr.ToString())
//	fmt.Println("Removed element:", removed)
//}
