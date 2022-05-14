package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    dummyNode := &ListNode{
        0,
        head,
    }

    prevNode := dummyNode
    currNode := head

    for currNode != nil {
        val := currNode.Val

        if currNode.Next != nil && currNode.Next.Val == val {
            // need start delete
            for currNode != nil && currNode.Val == val {
                prevNode.Next = currNode.Next
                currNode = currNode.Next
            }
        } else {
            prevNode = currNode
            currNode = currNode.Next
        }
    }

    return dummyNode.Next
}
