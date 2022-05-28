package leetcode

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    result := make([][]int, m, m)
    for i := 0; i < m; i++ {
        result[i] = make([]int, n, n)
    }

    for i := m-1; i >= 0; i-- {
        for j := n-1; j >=0; j-- {
            if i == m-1 && j == n-1 {
                result[i][j] = grid[i][j]
            } else if i == m-1 {
                result[i][j] = grid[i][j] + result[i][j+1]
            } else if j == n-1 {
                result[i][j] = grid[i][j] + result[i+1][j]
            } else {
                result[i][j] = grid[i][j] + min(result[i][j+1], result[i+1][j])
            }
        }
    }

    return result[0][0]
}

func min(i, j int) int {
    if i < j {
        return i
    } else {
        return j
    }
}
