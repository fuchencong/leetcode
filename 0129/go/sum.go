package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
    sum := 0

    var dfsSearch func(node *TreeNode, currVal int)
    dfsSearch = func(node *TreeNode, currVal int) {
        if node == nil {
            return
        }

        currVal = currVal * 10 + node.Val
        if node.Left == nil && node.Right == nil {
            sum += currVal
            return
        } else {
            if node.Left != nil {
                dfsSearch(node.Left, currVal)
            }
            if node.Right != nil {
                dfsSearch(node.Right, currVal)
            }
        }
    }

    dfsSearch(root, 0)
    return sum
}
