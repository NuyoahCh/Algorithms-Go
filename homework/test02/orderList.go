package main

/*
*
链表升序
一个链表，单索引是递增的，双索引是递减的，请对它进行升序排序，要求O(1)空间复杂度
个人感觉是合并双链表的变形题型
实现思路
1. 把一个链表拆分成为，奇偶两个链表
2. 把倒序的链表反转，通过正序实现
3. 将两个排好序的链表进行归并，也就是合并两个有序链表
*/
func orderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var odd, even *ListNode
	var tailA, tailB *ListNode
	cur := head
	index := 1

	for cur != nil {
		next := cur.Next
		cur.Next = nil

		if index%2 == 1 {
			if odd == nil {
				odd = cur
				tailA = cur
			} else {
				tailA.Next = cur
				tailA = cur
			}
		} else {
			if even == nil {
				even = cur
				tailB = cur
			} else {
				tailB.Next = cur
				tailB = cur
			}
		}

		cur = next
		index++
	}

	// 反转B链表（使其变为升序）
	even = reverseList1(even)

	// 合并两个升序链表
	return mergeTwoLists1(odd, even)
}

// 反转链表
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 合并两个升序链表
func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}
