package leetcode

func combine(n int, k int) [][]int {
    return dfsSearch(1, n, k)
}

func dfsSearch(min, max, k int) [][] int {
    result := [][]int{}

    if k == 0 {
        return result
    }

    for i := min; i <= max-k+1; i++ {
        tmpResults := dfsSearch(i+1, max, k-1)
        if len(tmpResults) == 0 {
            result = append(result, []int{i})
        } else {
            for j := 0; j < len(tmpResults); j++ {
                result = append(result, append([]int{i}, tmpResults[j]...))
            }
        }
    }

    return result
}
