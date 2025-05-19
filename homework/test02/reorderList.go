package main

/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L(0) → L(1) → … → L(n - 1) → L(n)

请将其重新排列后变为：
L(0) → L(n) → L(1) → L(n - 1) → L(2) → L(n - 2) → …
不能只是单纯地改变节点内部的值，而是需要实际的进行节点交换。
*/

// // DoublyListNode 建立双向链表
//
//	type DoublyListNode struct {
//		Value int
//		Pre   *DoublyListNode
//		Next  *DoublyListNode
//	}
//
// // 重排链表
//
//	func reorderList(head *DoublyListNode) *DoublyListNode {
//		dummy := &DoublyListNode{Value: 0, Pre: nil, Next: nil}
//		// 0 + n, 1 + (n - 1), 2 + (n - 2)
//		if head == nil || head.Next == nil {
//			return head
//		}
//		slow := head
//		fast := head
//		p := dummy
//		for slow != fast {
//			slow.Next =
//		}
//	}

// reorderList 重新排列链表，使其呈现交错顺序。
// 该函数不返回新的链表头节点，而是直接修改输入的链表。
func reorderList(head *ListNode) {
	// 先把所有节点装进栈里面，得到倒序结果
	type stack []*ListNode
	// 创建一个空的栈，用于存放链表中的所有节点
	var stk stack
	// p 是一个指针，初始化为链表头节点
	p := head
	// 把所有的元素都装到栈中去，得到一个倒序的访问方式
	for p != nil {
		// 将当前节点加入栈顶
		stk = append(stk, p)
		// 移动到下一个节点
		p = p.Next
	}

	// 重置 p 回到链表头
	p = head

	// 从栈中弹出节点，插入到当前 p 节点后面，实现链表的重排
	for p != nil {
		// 弹出栈顶的节点，即链表末尾的节点
		lastNode := stk[len(stk)-1]
		// 弹出后更新栈的内容
		stk = stk[:len(stk)-1]

		// 保存当前节点的下一个节点，后面还要用
		next := p.Next

		// 判断是否已经到了该停止的条件
		// 1. lastNode == next 说明链表节点数为偶数，中间已经相遇
		// 2. lastNode.Next == next 表示链表节点数为奇数，最后三个节点已重排完成
		if lastNode == next || lastNode.Next == next {
			// 结束条件，链表节点数为奇数或偶数时均适用
			// 防止形成环，确保尾节点指向 nil
			lastNode.Next = nil
			// 终止循环
			break
		}
		// 把尾部节点插入当前节点 p 的后面
		p.Next = lastNode
		// 尾部节点后面接上原来的 next
		lastNode.Next = next
		// 移动到 next 节点，继续处理后续部分
		p = next
	}
}
