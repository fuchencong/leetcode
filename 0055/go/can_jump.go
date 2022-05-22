package leetcode

func canJump(nums []int) bool {
    lenNums := len(nums)
    canJump := make([]bool, lenNums, lenNums)

    for i := lenNums-1; i >= 0; i-- {
        if i == lenNums-1 {
            canJump[i] = true
            continue
        }

        steps := nums[i]
        t := false
        for j := 0; j <= steps && i+j < lenNums; j++ {
            if canJump[i+j] {
                t = true
                break
            }
        }
        canJump[i] = t
    }

    return canJump[0]
}
