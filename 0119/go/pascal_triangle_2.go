package leetcode

func getRow(rowIndex int) []int {
    results := make([][]int, 0, rowIndex+1)

    results = append(results, []int{1})
    for i := 1; i <= rowIndex; i++ {
        currLevel := make([]int, 0, len(results[i-1])+1)
        for j := 0; j <= len(results[i-1]); j++ {
            if j == 0 {
                currLevel = append(currLevel, 1)
            } else if j == len(results[i-1]) {
                currLevel = append(currLevel, 1)
            } else {
                currLevel = append(currLevel, results[i-1][j-1] + results[i-1][j])
            }
        }
        results = append(results, currLevel)
    }

    return results[rowIndex]
}
