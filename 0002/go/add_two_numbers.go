package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	v := 0
	head := &ListNode{}
	prev := head

	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		n := &ListNode{
			Val: (v + v1 + v2) % 10,
		}
		v = (v + v1 + v2) / 10

		prev.Next = n
		prev = n
	}

	if v > 0 {
		n := &ListNode{
			Val: v % 10,
		}
		prev.Next = n
	}

	return head.Next
}
