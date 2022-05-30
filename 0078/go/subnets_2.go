package leetcode

func subsets(nums []int) [][]int {
    result := [][]int{}
    currSubnet := []int{}

    var dfsSearch func (nums []int)
    dfsSearch = func(nums []int ){
        if len(nums) == 0 {
            result = append(result, append([]int{}, currSubnet...))
            return
        }

        // choose nums[0]
        currSubnet = append(currSubnet, nums[0])
        dfsSearch(nums[1:])
        currSubnet = currSubnet[0:len(currSubnet)-1]

        // not choose nums[0]
        dfsSearch(nums[1:])

    }

    dfsSearch(nums)
    return result
}
