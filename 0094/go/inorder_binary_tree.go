package leetcode

 type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

func inorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    results := []int{}

    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        results = append(results, root.Val)
        root = root.Right
    }

    return results
}
