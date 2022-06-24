package leetcode

func isPalindrome(s string) bool {
    result := true
    bs := make([]byte, 0, len(s))

    for i := 0; i < len(s); i++ {
        b := s[i]
        if 'A' <= b && b <= 'Z' {
            b = 'a' + (b - 'A')
        }

        if ('a' <= b && b <= 'z') ||
           ('0' <= b && b <= '9') {
            bs = append(bs, b)
        }
    }

    for i, j := 0, len(bs) - 1; i < j; {
        if bs[i] != bs[j] {
            result = false
            break
        }

        i++
        j--
    }

    return result
}
