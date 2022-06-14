package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
    currLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}

    if root == nil {
        return 0
    }

    currDepth := 0
    reachLeaf := false
    for !reachLeaf  {
        currDepth++
        for _, node := range currLevel {
            if node.Left == nil && node.Right == nil {
                reachLeaf = true
                break
            }

            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }

        currLevel = nextLevel
        nextLevel = []*TreeNode{}
    }

    return currDepth
}
