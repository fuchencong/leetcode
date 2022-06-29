package leetcode

func wordBreak(s string, wordDict []string) bool {
    wordExist := make(map[string]bool)
    for _, word := range wordDict {
        wordExist[word] = true
    }

    n := len(s)
    searchResult := make([]bool, n, n)
    for i := n-1; i >= 0; i-- {
        for j := i+1; j <= n; j++ {
            if _, ok := wordExist[s[i:j]]; ok {
                if j >= n || searchResult[j] {
                    searchResult[i] = true
                }
            }
        }
    }

    return searchResult[0]
}
