package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfsSearch func(root*TreeNode, currSum int)
    result := false

    dfsSearch = func(node *TreeNode, currSum int) {
        if node == nil {
            return
        }

        currSum += node.Val
        if node.Left == nil && node.Right == nil && currSum == targetSum {
            result = true
            return
        }

        if !result && node.Left != nil {
            dfsSearch(node.Left, currSum)
        }

        if !result && node.Right != nil {
            dfsSearch(node.Right, currSum)
        }
    }

    dfsSearch(root, 0)
    return result
}
