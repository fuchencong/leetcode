package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    dummyNode := &TreeNode{}
    prevNode := dummyNode

    stack := []*TreeNode{}
    stack = append(stack, root)

    for len(stack) > 0 {
        root = stack[len(stack)-1]
        stack = stack[0:len(stack)-1]

        prevNode.Left = nil
        prevNode.Right = root

        if root.Right != nil {
            stack = append(stack, root.Right)
        }

        if root.Left != nil {
            stack = append(stack, root.Left)
        }

        prevNode = root
    }
}
