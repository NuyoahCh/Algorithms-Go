package main

/**
反转链表的迭代法
反转链表的递归法
*/

// ListNode 创建链表节点结构
type ListNode struct {
	Value int
	Next  *ListNode
}

// 反转链表的迭代法
func reverseList(head *ListNode) *ListNode {
	// 定义一个前驱指针 pre，初始为 nil，表示反转后链表的尾部
	var pre *ListNode
	// 定义一个当前指针 cur，初始指向链表头部 head
	cur := head
	// 遍历整个链表，直到 cur 为空
	for cur != nil {
		// 保存当前节点的下一个节点，防止断链
		next := cur.Next
		// 将当前节点的 next 指向前一个节点，实现指针反转
		cur.Next = pre
		// 前驱节点前移，当前节点成为新的 pre
		pre = cur
		// 当前节点后移，继续处理下一个节点
		cur = next
	}
	// 当 cur 为 nil 时，pre 指向新的头节点，返回它
	return pre
}

// 反转链表的递归法
func reverseListForMe(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseListForMe(head)
	// 当前节点的下一个节点的 next 指向当前节点（反转）
	head.Next.Next = head
	// 当前节点下节点断开，防止成环
	head.Next = nil
	return last
}
