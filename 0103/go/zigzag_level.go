package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    currLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}

    results := [][]int{}

    if root == nil {
        return results
    }

    reverseOrder := false
    for len(currLevel) != 0 {
        currLevelResults := []int{}
        for _, node := range currLevel {
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }

            if reverseOrder {
                currLevelResults = append([]int{node.Val}, currLevelResults...)
            } else {
                currLevelResults = append(currLevelResults, node.Val)
            }
        }

        results = append(results, currLevelResults)
        currLevel = nextLevel
        nextLevel = []*TreeNode{}
        reverseOrder = !reverseOrder
    }


    return results
}
