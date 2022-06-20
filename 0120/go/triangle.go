package leetcode

func minimumTotal(triangle [][]int) int {
    results := make([][]int, 0, len(triangle))

    results = append(results, []int{triangle[0][0]})

    for i := 1; i < len(triangle); i++ {
        currLevel := make([]int, 0, len(triangle[i]))
        for j := 0; j < len(triangle[i]); j++ {
            if j == 0 {
                currLevel = append(currLevel, results[i-1][j] + triangle[i][j])
            } else if j == len(triangle[i])-1 {
                currLevel = append(currLevel, results[i-1][j-1] + triangle[i][j])
            } else {
                currLevel = append(currLevel, min(results[i-1][j-1], results[i-1][j]) + triangle[i][j])
            }
        }
        results = append(results, currLevel)
    }

    minPathSum := results[len(results)-1][0]
    for i := 0; i < len(results[len(results) - 1]); i++ {
        if results[len(results)-1][i] < minPathSum {
            minPathSum = results[len(results)-1][i]
        }
    }

    return minPathSum
}


func min(i, j int) int {
    if i < j {
        return i
    } else {
        return j
    }
}
