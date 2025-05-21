package main

import "fmt"

/**
两个栈实现一个队列
*/

// MyQueue 定义 MyQueue 结构体
type MyQueue struct {
	storageStack   []int
	retrievalStack []int
}

// Constructor 构造函数
func Constructor() MyQueue {
	return MyQueue{
		storageStack:   make([]int, 0),
		retrievalStack: make([]int, 0),
	}
}

// Offer 将元素推到队列的末尾
func (mq *MyQueue) Offer(x int) {
	mq.storageStack = append(mq.storageStack, x)
}

// Poll 从队列的开头移除并返回元素
func (mq *MyQueue) Poll() int {
	if len(mq.retrievalStack) == 0 {
		mq.loadMore()
	}
	if len(mq.retrievalStack) == 0 {
		return -1
	}
	val := mq.retrievalStack[len(mq.retrievalStack)-1]
	mq.retrievalStack = mq.retrievalStack[:len(mq.retrievalStack)-1]
	return val
}

// Peek 返回队列开头的元素
func (mq *MyQueue) Peek() int {
	if len(mq.retrievalStack) == 0 {
		mq.loadMore()
	}
	if len(mq.retrievalStack) == 0 {
		return -1 // 假设队列为空时返回 -1
	}
	return mq.retrievalStack[len(mq.retrievalStack)-1]
}

// 判断队列是否为空
func (mq *MyQueue) IsEmpty() bool {
	return len(mq.storageStack) == 0 && len(mq.retrievalStack) == 0
}

func (mq *MyQueue) loadMore() {
	for len(mq.storageStack) > 0 {
		top := mq.storageStack[len(mq.storageStack)-1]
		mq.storageStack = mq.storageStack[:len(mq.storageStack)-1]
		mq.retrievalStack = append(mq.retrievalStack, top)
	}
}

// 测试示例
func main() {
	queue := Constructor()
	queue.Offer(1)
	queue.Offer(2)
	fmt.Println(queue.Peek())    // 输出 1
	fmt.Println(queue.Poll())    // 输出 1
	fmt.Println(queue.IsEmpty()) // 输出 false
}
