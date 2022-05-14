package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    var prevNode *ListNode
    currNode := head

    for currNode != nil {
        if prevNode != nil && prevNode.Val == currNode.Val {
            prevNode.Next = currNode.Next
        } else {
            prevNode = currNode
        }

        currNode = currNode.Next
    }

    return head
}
