package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if peer, ok := m[target-value]; ok {
			return []int{peer, index}
		}
		m[value] = index
	}
	return nil
}
