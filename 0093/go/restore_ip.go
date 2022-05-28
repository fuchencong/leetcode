package leetcode

import (
    "strings"
    "strconv"
)

func restoreIpAddresses(s string) []string {
    return splitIpAddress(s, 4, []string{})
}


func splitIpAddress(s string, nums int, currByteStrs []string) []string {
    results := []string{}

    if nums == 1 {
        if validIpByteStr(s) {
            currByteStrs = append(currByteStrs, s)
            if len(currByteStrs) == 4 {
                results = append(results, strings.Join(currByteStrs, "."))
            }
        }

        return results
    }

    for i :=  0; i < len(s) && i < 4; i++ {
        if validIpByteStr(s[0:i+1]) {
            currByteStrs = append(currByteStrs, s[0:i+1])
            results = append(results, splitIpAddress(s[i+1:], nums-1, currByteStrs)...)
            currByteStrs = currByteStrs[:len(currByteStrs)-1]
        }
    }

    return results
}

func validIpByteStr(s string) bool {
    if len(s) == 0 || len(s) > 3{
        return false
    }

    if s[0] == '0' && len(s) > 1 {
        return false
    }

    r, err := strconv.Atoi(s)
    if err != nil || r > 255 || r < 0 {
        return false
    }

    return true
}
