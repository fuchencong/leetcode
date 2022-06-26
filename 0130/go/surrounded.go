package leetcode

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])

    var dfsSearch func(r, c int)
    dfsSearch = func(r, c int) {
        if r < 0 || r >= m || c < 0 || c >= n {
            return
        }

        if board[r][c] != 'O' {return}

        board[r][c] = '$'
        dfsSearch(r+1, c)
        dfsSearch(r-1, c)
        dfsSearch(r, c+1)
        dfsSearch(r, c-1)
    }

    for r := 0; r < m; r++ {
        dfsSearch(r, 0)
        dfsSearch(r, n-1)
    }

    for c := 0; c < n; c++ {
        dfsSearch(0, c)
        dfsSearch(m-1, c)
    }

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if board[r][c] != '$' { board[r][c] = 'X' } else { board[r][c] = 'O' }
        }
    }



    return
}
