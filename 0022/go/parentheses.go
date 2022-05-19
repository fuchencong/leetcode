package leetcode

func generateParenthesis(n int) []string {
    const LeftParan = '('
    const RightParan = ')'

    if n == 0 {
        return nil
    }

    s := make([]byte, 2*n, 2*n)
    results := []string{}

    var dfsSearch func(currLeftCnt, currRightCnt int)
    dfsSearch = func(currLeftCnt, currRightCnt int) {
        if currLeftCnt + currRightCnt == n * 2 {
            results = append(results, string(s))
            return
        }

        if currLeftCnt == n {
            s[currLeftCnt+currRightCnt] = RightParan
            currRightCnt++
            dfsSearch(currLeftCnt, currRightCnt)
            currRightCnt--
        } else if currLeftCnt > currRightCnt {
            s[currLeftCnt+currRightCnt] = LeftParan
            currLeftCnt++
            dfsSearch(currLeftCnt, currRightCnt)
            currLeftCnt--

            s[currLeftCnt+currRightCnt] = RightParan
            currRightCnt++
            dfsSearch(currLeftCnt, currRightCnt)
            currRightCnt--

        } else if currLeftCnt == currRightCnt{
            s[currLeftCnt+currRightCnt] = LeftParan
            currLeftCnt++
            dfsSearch(currLeftCnt, currRightCnt)
            currLeftCnt--
        } else {
            panic("right > left")
        }
    }

    dfsSearch(0, 0)

    return results
}
