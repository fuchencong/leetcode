package leetcode

func findMin(nums []int) int {
    i := 0
    j := len(nums) - 1
    for i < j  {
        mid := i + (j - i)/2
        if nums[mid] > nums[j] {
            i = mid + 1
        } else if nums[mid] < nums[j] {
            j = mid
        } else {
            j--
        }
    }

    return nums[i]
}
