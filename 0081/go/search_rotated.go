package leetcode

func search(nums []int, target int) bool {
    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := (left+right)/2

        if nums[mid] == target {
            return true
        }

        // right is ordered
        if nums[mid] < nums[right] {
            if nums[mid] < target && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        // left is ordered
        } else if nums[mid] > nums[right] {
            if nums[left] <= target && target < nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            }
        } else {
            right--
        }
    }

    return false
}
