package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    results := make([][]int, m, m)
    for i := 0; i < m; i++ {
        results[i] = make([]int, n, n)
    }

    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if obstacleGrid[i][j] == 1 {
                results[i][j] = 0
                continue
            }

            if i == m-1 && j == n-1 {
                results[i][j] = 1
            } else if i == m-1 {
                results[i][j] = results[i][j+1]
            } else if j == n-1 {
                results[i][j] = results[i+1][j]
            } else {
                results[i][j] += results[i+1][j] + results[i][j+1]
            }
        }
    }

    return results[0][0]
}
