package leetcode

func permute(nums []int) [][]int {
    choosed := make([]bool, len(nums), len(nums))
    return _permute(nums, choosed)
}

func _permute(nums []int, choosed []bool) [][]int{
    results := [][]int{}

    for i := 0; i < len(nums); i++ {
        if choosed[i] {
            continue
        }

        choosed[i] = true
        subResults := _permute(nums, choosed)
        choosed[i] = false

        if len(subResults) == 0 {
            results = append(results, []int{nums[i]})
        } else {
            for _, s := range(subResults) {
                r := make([]int, 0, len(nums))
                r = append(r, nums[i])
                t := append(r, s...)

                results = append(results, t)
            }
        }
    }

    return results
}
