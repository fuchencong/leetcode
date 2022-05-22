package leetcode

func spiralOrder(matrix [][]int) []int {
    rowsNum := len(matrix)
    colsNum := len(matrix[0])

    leftBorder := 0
    rightBorder := colsNum-1
    upBorder := 1
    downBorder := rowsNum-1

    result := make([]int, 0, rowsNum*colsNum)
    var goFunc [4]func(matrix [][]int, currRow, currCol int) (int, int)

    goFunc[0] = func(matrix [][]int, currRow, currCol int) (int, int){
        for currCol <= rightBorder {
            result = append(result, matrix[currRow][currCol])
            currCol++
        }
        rightBorder--
        return currRow+1, currCol-1
    }

    goFunc[1] = func(matrix [][]int, currRow, currCol int) (int, int){
        for currRow < downBorder {
            result = append(result, matrix[currRow][currCol])
            currRow++
        }
        downBorder--
        return currRow, currCol
    }

    goFunc[2] = func(matrix [][]int, currRow, currCol int) (int, int){
        for currCol >= leftBorder {
            result = append(result, matrix[currRow][currCol])
            currCol--
        }
        leftBorder++
        return currRow-1, currCol+1
    }

    goFunc[3] = func(matrix [][]int, currRow, currCol int) (int, int){
        for currRow > upBorder {
            result = append(result, matrix[currRow][currCol])
            currRow--
        }
        upBorder++
        return currRow, currCol
    }



    i, currRow, currCol := 0, 0, 0
    for true {
        i = i % len(goFunc)
        currRow, currCol = goFunc[i](matrix, currRow, currCol)

        if len(result) == rowsNum * colsNum {
            break
        }
        i++
    }

    return result
}
