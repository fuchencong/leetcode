package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func buildTree(preorder []int, inorder []int) *TreeNode {
    root := &TreeNode {
        Val: preorder[0],
    }

    leftLen, rightLen := 0, 0
    for i, v := range inorder {
        if v == preorder[0] {
            leftLen = i
        }
    }
    rightLen = len(preorder) - leftLen - 1

    if leftLen > 0 {
        root.Left = buildTree(preorder[1:1+leftLen], inorder[0:leftLen])
    }

    if rightLen > 0 {
        root.Right = buildTree(preorder[1+leftLen:], inorder[leftLen+1:])
    }

    return root
}
