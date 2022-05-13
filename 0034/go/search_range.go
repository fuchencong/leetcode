package leetcode

func searchRange(nums []int, target int) []int {
    startPos, endPos := -1, -1

    left := 0
    right := len(nums) - 1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            startPos = binarySearch(nums, target, left, mid, true)
            endPos = binarySearch(nums, target, mid, right, false)
            break
        }
    }

    return []int{startPos, endPos}
}

func binarySearch(nums []int, target int, left, right int, getFirst bool) int {
    targetPos := -1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            targetPos = mid
            if getFirst {
                // get first target
                right = mid - 1
            } else {
                // get last target
                left = mid + 1
            }
        }
    }

    return targetPos
}

