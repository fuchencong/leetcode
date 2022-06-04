package leetcode

func grayCode(n int) []int {
    if n == 1 {
        return []int{0, 1}
    }

    results := grayCode(n-1)
    for i := len(results)-1; i >=0 ; i-- {
        results = append(results, (results[i] | (1 << (n-1))))
    }

    return results
}
