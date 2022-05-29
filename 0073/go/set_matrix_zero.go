package leetcode

func setZeroes(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])

    zeroFirstRow, zeroFirstColumn := false, false

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                if i == 0 {
                    zeroFirstRow = true
                }
                if j == 0 {
                    zeroFirstColumn = true
                }

                matrix[i][0] = 0
                matrix[0][j] = 0
            }

        }
    }

    for i := m-1; i >= 1; i-- {
        if matrix[i][0] == 0 {
            for j := n-1; j >= 1; j-- {
                matrix[i][j] = 0
            }
        }
    }

    for j := n-1; j >= 1; j-- {
        if matrix[0][j] == 0 {
            for i := m-1; i >= 1; i-- {
                matrix[i][j] = 0
            }
        }
    }

    if zeroFirstColumn {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }

    if zeroFirstRow {
        for i := 0; i < n; i++ {
            matrix[0][i] = 0
        }

    }

    return
}
