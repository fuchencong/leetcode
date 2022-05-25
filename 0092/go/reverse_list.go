package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

    if left == right {
        return head
    }

    dummyNode := &ListNode{0, head}
    prevNode := dummyNode
    currNode := head
    cnt := 1

    var leftPrev, leftNode, rightPrev, rightNode *ListNode
    for currNode != nil {
        nextNode := currNode.Next

        if cnt >= left && cnt <= right {
            if cnt == left {
               leftPrev = prevNode
               leftNode = currNode
            } else if cnt == right {
                rightPrev = prevNode
                rightNode = currNode
            } else {
                currNode.Next = prevNode
            }
        }

        prevNode = currNode
        currNode = nextNode
        cnt++
    }

    leftPrev.Next=rightNode

    leftNode.Next, rightNode.Next = rightNode.Next, rightPrev

    return dummyNode.Next
}
