package leetcode

import (
    "sort"
)


func threeSumClosest(nums []int, target int) int {
    length := len(nums)

    if length < 3 {
        panic("less 3 numbers")
    }

    sort.Ints(nums)

    var result, minDiff int
    first := true
    for i := 0; i < length; i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }

        left := i+1
        right := length-1

        for left < right {
            sum := nums[left] + nums[right] + nums[i]
            currDiff := abs(sum - target)
            if first || currDiff < minDiff {
                minDiff = currDiff
                result = sum
                first = false
            }

            if sum > target {
                right--
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else {
                left++
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
            }
        }
    }

    return result
}

func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}
