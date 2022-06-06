package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)

    results := [][]int{}
    subnet := []int{}

    var dfsSearch func(bool, int)
    dfsSearch = func(preChoose bool, index int) {
        if index == n {
            results = append(results, append([]int{}, subnet...))
            return
        }

        dfsSearch(false, index+1)

        if !preChoose && index > 0 && nums[index-1] == nums[index] {
            return
        }

        subnet = append(subnet, nums[index])
        dfsSearch(true, index+1)
        subnet = subnet[0:len(subnet)-1]
    }

    dfsSearch(false, 0)

    return results
}
