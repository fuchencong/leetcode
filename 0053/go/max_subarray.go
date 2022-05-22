package leetcode

func maxSubArray(nums []int) int {
    lenNums := len(nums)

    if lenNums == 0 {
        return 0
    }

    maxSum := nums[0]
    for i := 0; i < lenNums; i++ {
        currSum := nums[i]
        maxSum = max(maxSum, currSum)

        j := i+1
        for currSum >=0 && j < lenNums {
            currSum += nums[j]
            maxSum = max(maxSum, currSum)
            j++
        }
    }

    return maxSum
}

func max(i, j int) int {
    if i >= j {
        return i
    } else {
        return j
    }
}
