package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    currLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}

    results := [][]int{}

    if root == nil {
        return results;
    }

    for len(currLevel) > 0 {
        currResult := []int{}
        for _, node := range currLevel {
            currResult = append(currResult, node.Val)
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }

        results = append(results, currResult)
        currLevel = nextLevel
        nextLevel = []*TreeNode{}
    }

    return results
}
