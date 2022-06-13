package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    lenNums := len(nums)
    if lenNums == 0 {
        return nil
    }

    mid := lenNums/2
    root := &TreeNode{
        nums[mid],
        nil,
        nil,
    }

    if mid > 0 {
        root.Left = sortedArrayToBST(nums[0:mid])
    }

    if mid + 1 < lenNums {
        root.Right = sortedArrayToBST(nums[mid+1:lenNums])
    }

    return root
}
