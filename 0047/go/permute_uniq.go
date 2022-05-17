package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
    used := make([]bool, len(nums), len(nums))

    results := [][]int{}

    permutation := []int{}
    sort.Ints(nums)
    var dfsPermute func(int)

    dfsPermute = func (index int) {
        if index == len(nums) {
            results = append(results, append([]int(nil), permutation...))
            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
                continue
            }


            used[i] = true
            permutation = append(permutation, nums[i])
            dfsPermute(index+1)
            used[i] = false
            permutation = permutation[0:index]

        }
    }

    dfsPermute(0)

    return results
}
