package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    result := 0

    var dfsSearch func(root *TreeNode, currResult int)
    dfsSearch = func(root *TreeNode, currResult int) {
        if root == nil {
            if currResult > result {
                result = currResult
            }
            return
        }

        currResult++
        dfsSearch(root.Left, currResult)
        dfsSearch(root.Right, currResult)
    }

    dfsSearch(root, 0)
    return result
}
