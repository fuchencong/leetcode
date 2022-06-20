package leetcode

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    root.Next = nil
    currLevelHead := root

    for currLevelHead != nil {
        prevNode := (*Node)(nil)
        nextLevelHead := (*Node)(nil)

        for currLevelHead != nil {
            if currLevelHead.Left != nil {
                if prevNode != nil {
                    prevNode.Next = currLevelHead.Left
                }
                currLevelHead.Left.Next = nil
                prevNode = currLevelHead.Left

                // set head of next level list
                if nextLevelHead == nil {
                    nextLevelHead = currLevelHead.Left
                }
            }

            if currLevelHead.Right != nil {
                if prevNode != nil {
                    prevNode.Next = currLevelHead.Right
                }
                currLevelHead.Right.Next = nil
                prevNode = currLevelHead.Right

                // set head of next level list
                if nextLevelHead == nil {
                    nextLevelHead = currLevelHead.Right
                }
            }

            currLevelHead = currLevelHead.Next
        }

        currLevelHead = nextLevelHead
    }

    return root
}

