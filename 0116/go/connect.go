package leetcode

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
    if root == nil {
        return root
    }

    root.Next = nil

    currLevelHead := root;
    var nextLevelHead *Node

    for currLevelHead != nil {
        nextLevelHead = currLevelHead.Left
        if nextLevelHead == nil {
            break
        }

        prevNode := (*Node)(nil)
        for currLevelHead != nil {
            if prevNode != nil {
                prevNode.Next = currLevelHead.Left
            }
            currLevelHead.Left.Next = currLevelHead.Right
            currLevelHead.Right.Next = nil

            prevNode = currLevelHead.Right
            currLevelHead = currLevelHead.Next
        }

        currLevelHead = nextLevelHead
    }

    return root
}
