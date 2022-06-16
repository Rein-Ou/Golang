package main

type ListNode struct {
	Val  int
	next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.next = &ListNode{Val: carry}
			tail = tail.next
		}
	}
	if carry > 0 {
		tail.next = &ListNode{Val: carry}
	}
	return
}
