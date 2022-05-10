package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode{
    left  := head
    right := head

    for i := 0; i < n - 1; i ++ {
        right = right.Next
    }

    var prev *ListNode
    for right.Next != nil {
        prev = left
        right = right.Next
        left = left.Next
    }

    if left == head {
        return head.Next
    } else {
        prev.Next = left.Next
        return head
    }
}
