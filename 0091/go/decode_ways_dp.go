package leetcode

func numDecodingsDP(s string) int {
    n := len(s)
    results := make([]int, n+1, n+1)
    results[0] = 1

    for i := 0; i < len(s); i++ {
        if s[i] != '0' {
            results[i+1] += results[i]
        }

        if i >= 1 && s[i-1] != '0' && ((s[i-1]-'0')*10 + (s[i]-'0') <= 26) {
            results[i+1] += results[i-1]
        }
    }

    return results[n]
}
