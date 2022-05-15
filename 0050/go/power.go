package leetcode

func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1 / pow(x, -n)
    }

    return pow(x, n)
}

func pow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }

    if n == 1 {
        return x
    }

    t := pow(x, n / 2)
    result := t * t
    if n & 1 != 0 {
        result = result * x
    }

    return result
}
