package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }

    rootVal := postorder[len(postorder)-1]
    leftLen := 0
    for inorder[leftLen] != rootVal {
        leftLen++
    }

    root := &TreeNode{
        rootVal,
        nil,
        nil,
    }

    root.Left = buildTree(inorder[0:leftLen], postorder[0:leftLen])
    root.Right = buildTree(inorder[leftLen+1:], postorder[leftLen:len(postorder)-1])
    return root
}

