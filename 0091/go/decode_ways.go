package leetcode

func numDecodings(s string) int {
    results := 0

    if len(s) == 0 {
        return 0
    }

    var dfsSearch func(string)
    dfsSearch = func(ss string) {
        if len(ss) == 0 {
            results++
            return
        }

        for i := 0; i < len(ss) && i < 2; i++ {
            if validStr(ss[:i+1]) {
                dfsSearch(ss[i+1:])
            }
        }
    }

    dfsSearch(s)
    return results
}

func validStr(s string) bool {
    if len(s) == 0 || len(s) > 2 {
        return false;
    }

    if s[0] == '0' {
        return false
    }

    if (len(s) == 2 && (s[0] > '2' || (s[0] == '2' && s[1] > '6'))) {
        return false
    }

    return true
}

