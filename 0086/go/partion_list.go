package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}


func partition(head *ListNode, x int) *ListNode {
    var insertNode, insertPrevNode *ListNode

    dummyNode := &ListNode {
        0,
        head,
    }

    prevNode := dummyNode
    currNode := head

    for currNode != nil {
        if currNode.Val >= x {
            if insertNode == nil {
                insertNode = currNode
                insertPrevNode = prevNode
            }

            prevNode = currNode
            currNode = currNode.Next
        } else if currNode.Val < x {
            if insertNode != nil {
                // remove
                prevNode.Next = currNode.Next

                // insert
                insertPrevNode.Next = currNode
                currNode.Next = insertNode
                insertPrevNode = currNode

                currNode = prevNode.Next
            } else {
                prevNode = currNode
                currNode = currNode.Next
            }
        }
    }

    return dummyNode.Next
}
