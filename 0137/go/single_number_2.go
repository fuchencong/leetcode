package leetcode

func singleNumber(nums []int) int {
    counts := [32]int{}

    var result int32 = 0
    for i := 0; i < 32; i++ {
        for _, n := range nums {
            counts[i] += (n >> i) & 1
        }

        if counts[i] % 3 != 0 {
            result += (1 << i)
        }
    }

    return (int)(result)
}
