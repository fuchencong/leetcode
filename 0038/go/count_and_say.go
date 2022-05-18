package leetcode

import "fmt"

func countAndSay(n int) string {
    results := make([]string, n, n)

    for i := 0; i < n; i++ {
        if i == 0 {
            results[i] = "1"
            continue
        }

        results[i] = func(s string) string {
            rs := ""

            for j := 0; j < len(s); {
                count := 1
                for j + 1 < len(s) && s[j] == s[j+1] {
                    j++
                    count++
                }

                ss := fmt.Sprintf("%d%s", count, string(s[j]))
                j++
                rs += ss
            }

            return rs
        }(results[i-1])
    }

    return results[n-1]
}
