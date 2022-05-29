package leetcode

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])

    var indexToPos func(index int) (x, y int)
    indexToPos = func(index int) (x, y int) {
        x = index / n
        y = index % n
        return
    }

    left, right := 0, m*n-1
    for left <= right {
        mid := (left + right)/2
        midX, midY := indexToPos(mid)

        if matrix[midX][midY] == target {
            return  true
        } else if matrix[midX][midY] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return false
}

