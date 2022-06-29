package leetcode

func wordBreakDfs(s string, wordDict []string) bool {
    wordExist := make(map[string]bool)
    for _, word := range wordDict {
        wordExist[word] = true
    }

    var dfsSearch func(s string) bool
    dfsSearch = func(s string) bool {
        if len(s) == 0 {
            return true
        }

        for i := 0; i < len(s); i++ {
            if _, ok := wordExist[s[0:i+1]]; ok {
                if i+1 >= len(s) || dfsSearch(s[i+1:]) {
                    return true
                }
            }
        }

        return false
    }

    return dfsSearch(s)
}
