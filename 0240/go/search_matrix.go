package leetcode

func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    if rows == 0 {
        return false
    }

    columns := len(matrix[0])
    if columns == 0 {
        return false
    }

    i := 0
    j := columns - 1
    for i < rows &&  j >= 0 {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] > target {
            j--
        } else {
            i++
        }
    }

    return false
}
