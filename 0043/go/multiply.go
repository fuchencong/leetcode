package leetcode

func multiply(num1 string, num2 string) string {
    len1 := len(num1)
    len2 := len(num2)

    result := make([]int, len1 + len2, len1 + len2)

    for i := 0; i < len1; i++ {
        for j := 0; j < len2; j++ {
            t1 := int(num1[len1-1-i]-'0')
            t2 := int(num2[len2-1-j]-'0')

            currPos := i + j
            result[currPos] += (t1 * t2) % 10
            result[currPos+1] += (t1 * t2) / 10
        }
    }

    s := make([]byte, 0, len1+len2)
    for i := 0; i < len1+len2; i++ {
        if result[i] >= 10 {
            s = append(s, byte('0' + result[i] % 10))
            result[i+1] += result[i] / 10
        } else {
            s = append(s, byte('0' + result[i]))
        }
    }

    ss := []byte{}
    skipZero := true
    for i := len(s) - 1; i >=0; i-- {
        if s[i] == '0' && skipZero {
            continue
        }

        ss = append(ss, s[i])
        skipZero = false
    }

    if len(ss) == 0 {
        ss = append(ss, '0')
    }

    return string(ss)
}
