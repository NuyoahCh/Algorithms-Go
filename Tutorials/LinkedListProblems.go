package main

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	size := getListSize(head)
	step := (size - 1) / 2
	cur := head
	for i := 0; i < step; i++ {
		cur = cur.Next
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return cur
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	cur1 := head1
	cur2 := head2
	for cur1 != nil && cur2 != nil {
		if cur1.Value < cur2.Value {
			dummy.Next = cur1
			cur1 = cur1.Next
		} else {
			dummy.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func cycleEntrance(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func getListSize(head *ListNode) int {
	count := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}
	return count
}
