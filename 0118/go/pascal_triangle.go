package leetcode

func generate(numRows int) [][]int {
    results := [][]int{}

    results = append(results, []int{1})
    for i := 1; i < numRows; i++ {
        currLevel := []int{}
        for j := 0; j < len(results[i-1]) + 1; j++ {
            if j == 0 {
                currLevel = append(currLevel, 1)
            } else if j == len(results[i-1]) {
                currLevel  = append(currLevel, 1)
            } else {
                currLevel = append(currLevel, results[i-1][j-1]+results[i-1][j])
            }
        }
        results = append(results, currLevel)
    }

    return results
}
