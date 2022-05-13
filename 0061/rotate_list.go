package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

    listLen := 0
    t := head
    tail := head

    for t != nil {
        tail = t
        listLen++
        t = t.Next
    }

    if listLen == 0 || k == 0 {
        return head
    }

    k = k % listLen
    if k == 0 {
        return head
    }

    targetNode := head
    for i := 0; i < listLen - k - 1; i++ {
        targetNode = targetNode.Next
    }

    newHead := targetNode.Next
    targetNode.Next = nil
    tail.Next = head

    return newHead
}
