package leetcode

func lengthOfLastWord(s string) int {
    wordStart := false

    result, currCnt := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            if !wordStart {
                wordStart = true
            }

            if wordStart {
                currCnt++
            }

        } else {
            if wordStart {
                wordStart = false
                result = currCnt
                currCnt = 0
            }
        }
    }

    if wordStart {
        result = currCnt
    }

    return result
}
