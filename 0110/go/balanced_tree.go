package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    result := true

    var dfsSearch func(root *TreeNode, currLevel int) int
    dfsSearch = func(root *TreeNode, currLevel int) int {
        if root == nil {
            return currLevel
        }

        currLevel++
        leftLevel := dfsSearch(root.Left, currLevel)
        rightLevel := dfsSearch(root.Right, currLevel)

        if leftLevel - rightLevel > 1 ||
           rightLevel - leftLevel > 1 {
            result = false
        }

        return max(leftLevel, rightLevel)
    }

    dfsSearch(root, 0)
    return result
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
