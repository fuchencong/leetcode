package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
    results := [][]int{}

    var dfsSearch func(currLevel []*TreeNode)
    dfsSearch = func(currLevel []*TreeNode) {
        if len(currLevel) == 0 {
            return
        }

        currResult := make([]int, 0, len(currLevel))
        nextLevel := []*TreeNode{}
        for _, node := range currLevel {
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
            currResult = append(currResult, node.Val)
        }

        dfsSearch(nextLevel)
        results = append(results, currResult)
    }

    if root != nil {
        dfsSearch([]*TreeNode{root})
    }

    return results
}
