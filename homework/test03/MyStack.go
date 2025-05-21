package main

import (
	"container/list"
)

/**
用一个或两个队列实现一个后入先出（LIFO）的栈
*/

type MyStack struct {
	queue *list.List
}

func Constructor1() MyStack {
	return MyStack{
		queue: list.New(),
	}
}

func (ms *MyStack) Pop() int {
	if ms.queue.Len() == 0 {
		return -1
	}
	return ms.queue.Remove(ms.queue.Front()).(int)
}

// Top 返回栈顶元素
func (ms *MyStack) Top() int {
	if ms.queue.Len() == 0 {
		return -1 // 假设栈为空时返回 -1
	}
	return ms.queue.Front().Value.(int)
}

// Empty 判断栈是否为空
func (ms *MyStack) Empty() bool {
	return ms.queue.Len() == 0
}
