package leetcode

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    var result int64
    t := x
    for t != 0 {
        result = result * 10 + int64(t % 10)
        t = t / 10
    }

    return result==int64(x)
}

