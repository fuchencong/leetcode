package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    nums := []int{}

    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }

    var constructBST func(nums []int, left, right int) *TreeNode
    constructBST = func(nums []int, left, right int) *TreeNode {
        if left > right {
            return nil
        }

        mid := (left + right)/2
        root := &TreeNode{nums[mid], nil, nil}
        root.Left = constructBST(nums, left, mid-1)
        root.Right = constructBST(nums, mid+1, right)

        return root
    }

   return constructBST(nums, 0, len(nums)-1)
}
