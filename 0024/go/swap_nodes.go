package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    dummyNode := &ListNode{0, head}
    prevNode := dummyNode
    currNode := head

    for currNode != nil {
        nextNode := currNode.Next
        if nextNode == nil {
            break
        }

        prevNode.Next = nextNode
        currNode.Next = nextNode.Next
        nextNode.Next = currNode

        prevNode = currNode
        currNode = currNode.Next
    }

    return dummyNode.Next
}
