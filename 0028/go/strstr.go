package leetcode

func strStr(haystack string, needle string) int {
    result := -1

    lenH := len(haystack)
    lenN := len(needle)

    if lenN == 0 {
        return 0
    } else if lenH < lenN {
        return result
    }

    for i := 0; i <= lenH - lenN; i++ {
        match := true
        for j := 0; j < lenN; j++ {
            if haystack[i + j] != needle[j] {
                match = false
                break
            }
        }

        if match {
            result = i
            break
        }
    }

    return result
}
