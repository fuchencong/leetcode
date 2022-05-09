package leetcode

func exist(board [][]byte, word string) bool {
    rows := len(board)
    columns := len(board[0])
    used := make([][]bool, rows, rows)
    for i := 0; i < rows; i++ {
        used[i] = make([]bool, columns, columns)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if searchFromPos(board, used, word, i, j) {
                return true
            }
        }
    }

    return false
}

func searchFromPos(board[][]byte, used[][]bool, word string, i, j int) bool {
    rows := len(board)
    columns := len(board[0])

    if board[i][j] == word[0] && !used[i][j] {
        used[i][j] = true

        if len(word) <= 1 {
            return true
        }

        if i > 0 && searchFromPos(board, used, word[1:], i - 1, j) {
            return true
        }

        if i < rows-1 && searchFromPos(board, used, word[1:], i+1, j) {
            return true
        }

        if j > 0 && searchFromPos(board, used, word[1:], i, j - 1) {
            return true
        }

        if j < columns-1 && searchFromPos(board, used, word[1:], i, j + 1) {
            return true
        }

        used[i][j] = false
    }

    return false
}
