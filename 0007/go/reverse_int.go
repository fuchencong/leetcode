package leetcode

import "math"

func reverse(x int) int {
    res := 0

    for x != 0 {
        i := x % 10
        if res > 0 && ((res > math.MaxInt32 / 10) || res == math.MaxInt32 / 10 && i > 7){
            return 0
        }
        if res < 0 && ((res < math.MinInt32 / 10) || res == math.MinInt32 / 10 && i < -8){
            return 0
        }

        res = res * 10 + i
        x = x / 10
    }

    return res
}
