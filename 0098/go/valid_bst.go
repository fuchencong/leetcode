package leetcode

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}

    prevValue := math.MinInt64
    for root !=nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        root = stack[len(stack)-1]
        stack = stack[0:len(stack)-1]

        if root.Val <= prevValue {
            return false
        }

        prevValue = root.Val
        root = root.Right
    }

    return true
}
