package leetcode

func mySqrt(x int) int {
    if x == 0 {
        return 0
    } else if x < 4 {
        // use x/2 just right when x>=4
        return 1
    }

    left := 0
    right := x
    for left <= right {
        mid := (left + right) / 2
        v := mid * mid
        if v == x {
            return mid
        } else if v > x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return right
}
