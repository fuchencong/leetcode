package leetcode

func addBinary(a string, b string) string {
    result := []byte{}

    la := len(a)
    lb := len(b)

    var extra byte
    for i := 0; i < la || i < lb; i++ {
        var ba, bb byte

        if i >= la {
            ba = 0
            bb = b[lb - 1 - i] - '0'
        } else if i >= lb {
            ba = a[la - 1 - i] - '0'
            bb = 0
        } else {
            ba = a[la - 1 - i] - '0'
            bb = b[lb - 1 - i] - '0'
        }

        result = append(result, '0' + (extra + ba + bb) % 2)
        extra = (extra + ba + bb) / 2
    }

    if extra != 0 {
        result = append(result, '0' + extra)
    }

    // reverse byte
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }

    return string(result)
}
