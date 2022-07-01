package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head.Next

    result := false
    for fast != nil && fast.Next != nil {
        if slow == fast {
            result = true
            break
        }

        slow = slow.Next
        fast = fast.Next.Next
    }

    return result
}
