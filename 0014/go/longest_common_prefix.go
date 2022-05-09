package leetcode

func longestCommonPrefix(strs []string) string {
    index := 0
    end := false

    for true {
        for i := 0; i < len(strs); i++ {
            if index >= len(strs[i]) {
                end = true
                break
            }

            if strs[i][index] != strs[0][index] {
                end = true
                break
            }
        }

        if end {
            break
        }

        index++
    }

    return strs[0][0:index]
}
