package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyNode := &ListNode{}
    currNode := dummyNode

    for list1 != nil || list2 != nil {
        if list1 == nil {
            currNode.Next = list2
            break
        }

        if list2 == nil {
            currNode.Next = list1
            break
        }

        if list1.Val < list2.Val {
            currNode.Next = list1
            list1 = list1.Next
        } else {
            currNode.Next = list2
            list2 = list2.Next
        }

        currNode = currNode.Next
    }

    return dummyNode.Next
}
