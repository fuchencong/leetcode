package leetcode

func longestPalindrome(s string) string {
    resLen, resLeft, resRight := 0, 0, 0

    calcPalindrome := func(i, j int) {
        for i >= 0 && j < len(s) && s[i] == s[j] {
            i--
            j++
        }

        left := i+1
        right := j-1
        if right - left + 1 > resLen {
            resLeft = left
            resRight = right
            resLen = right - left + 1
        }
    }

    for i := 0; i < len(s); i++ {
        j := i
        calcPalindrome(i, j)
        j = i+1
        calcPalindrome(i, j)
    }
    return s[resLeft:resRight+1]
}

