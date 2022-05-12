package leetcode

import "sort"

func threeSum(nums []int) [][]int {
    results := make([][]int, 0, 1000)

    length := len(nums)
    if length < 3 {
        return results
    }

    sort.Ints(nums)
    for i := 0; i < length; i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        target := -nums[i]
        left := i+1
        right := length-1

        for left < right {
            if nums[left] + nums[right] < target {
                left++
            } else if nums[left] + nums[right] > target {
                right--
            } else {
                results = append(results, []int{nums[i], nums[left], nums[right]})
                left++
                right--

                for left < right && nums[left] == nums[left-1] {
                    left++
                }

                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            }
        }
    }

    return results
}
