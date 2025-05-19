package main

/**
快慢指针定位节点
计数法定位节点
*/

// 快慢指针定位节点
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 计数法定位节点
func middleNodeI(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	size := getListSize(head)
	// index
	step := (size - 1) / 2
	cur := head
	for i := 0; i < step; i++ {
		cur = cur.Next
	}
	return cur
}

// 获取链表长度
func getListSize(head *ListNode) int {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}
