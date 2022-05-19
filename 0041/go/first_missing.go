package leetcode

func firstMissingPositive(nums []int) int {
    lenNums := len(nums)

    for i := 0; i < lenNums; i++ {
        value := nums[i]
        for (1 <= nums[i] && nums[i] <= lenNums) && i + 1 != value && value != nums[value-1] {
            nums[i], nums[value-1] = nums[value-1], nums[i]
            value = nums[i]
        }
    }

    for i := 0; i < lenNums; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }

    return lenNums+1
}
