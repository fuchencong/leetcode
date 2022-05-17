package leetcode

func jump(nums []int) int {
    minSteps := make([]int, len(nums), len(nums))
    endPos := len(nums) - 1
    const UnreachableSteps= (10000 + 1)

    for i := endPos; i >= 0; i-- {
        if i == endPos {
            minSteps[i] = 0
            continue
        }

        minSteps[i] = UnreachableSteps
        for j := 1; j <= nums[i]; j++ {
            if i + j >= endPos {
                minSteps[i] = 1
            } else if 1 + minSteps[i+j] < minSteps[i] {
                minSteps[i] = 1+minSteps[i+j]
            }
        }
    }

    return minSteps[0]
}
