package leetcode 
//package leetcode
//
//func subsets(nums []int) [][]int {
//    result := [][]int{}
//    currSubnet := []int{}
//
//    var dfsSearch func (index int)
//    dfsSearch = func(index int){
//        if len(nums) == index {
//            result = append(result, append([]int{}, currSubnet...))
//            return
//        }
//
//        // choose nums[index]
//        currSubnet = append(currSubnet, nums[index])
//        dfsSearch(index+1)
//        currSubnet = currSubnet[0:len(currSubnet)-1]
//
//        // not choose nums[0]
//        dfsSearch(index+1)
//
//    }
//
//    dfsSearch(0)
//    return result
//}
