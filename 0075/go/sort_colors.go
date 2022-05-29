package leetcode

func sortColors(nums []int)  {
    colorsCnt := []int{0, 0, 0}

    for i := 0; i < len(nums); i++ {
        colorsCnt[nums[i]]++
    }

    currPos := 0
    for i := 0; i < len(colorsCnt); i++ {
        for j := 0; j < colorsCnt[i]; j++ {
            nums[currPos] = i
            currPos++
        }
    }

    return
}
