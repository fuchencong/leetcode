package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    results := [][]int{}

    var dfsSearch func(node *TreeNode, currSum int, currNodes []int)
    dfsSearch = func(node *TreeNode, currSum int, currNodes []int) {
        if node == nil {
            return
        }

        currSum += node.Val
        currNodes = append(currNodes, node.Val)
        if node.Left == nil && node.Right == nil && currSum == targetSum {
            results = append(results, append([]int{}, currNodes...))
            return
        }

        if node.Left != nil {
            dfsSearch(node.Left, currSum, currNodes)
        }

        if node.Right != nil {
            dfsSearch(node.Right, currSum, currNodes)
        }

        currSum -= node.Val
        currNodes = currNodes[0:len(currNodes)-1]
    }

    dfsSearch(root, 0, []int{})
    return results
}
