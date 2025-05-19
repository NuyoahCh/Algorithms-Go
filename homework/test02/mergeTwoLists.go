package main

/**
核心就是创建虚拟头节点，把我们比较判断完成的值放置到dummy中
虚拟头节点的使用总结：
当我们真的需要使用开辟一个新的空间，放置之前的元素，如果直接在之前的结构上进行操作复杂
*/

// 合并链表
func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p1, p2 := head1, head2
	p := dummy
	for p1 != p2 {
		if p1.Value < p2.Value {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p = p2.Next
	}
	if p2 == nil {
		p = p1.Next
	}
	return dummy.Next
}
