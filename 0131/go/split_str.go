package leetcode

func partition(s string) [][]string {
    n := len(s)

    isPalindrome := make([][]bool, n, n)
    for i := 0; i < n; i++ {
        isPalindrome[i] = make([]bool, n, n)
    }

    for i := n-1; i >= 0; i-- {
        for j := 0; j < n; j++ {
            if i >= j {
                isPalindrome[i][j] = true
            } else {
                isPalindrome[i][j] = (s[i]==s[j] && isPalindrome[i+1][j-1])
            }
        }
    }

    results := [][]string{}
    currSplits := []string{}

    var dfsSearch func (int)
    dfsSearch = func(start int) {
        if start == n {
            results = append(results, append([]string{}, currSplits...))
            return
        }

        for i := start; i < n; i++ {
            if isPalindrome[start][i] {
                currSplits = append(currSplits, s[start:i+1])
                dfsSearch(i+1)
                currSplits = currSplits[0:len(currSplits)-1]
            }
        }
    }

    dfsSearch(0)
    return results
}
