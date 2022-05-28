package leetcode

func plusOne(digits []int) []int {
    extra := 0

    for i := len(digits) - 1; i >= 0; i-- {
        if i == len(digits) - 1 {
            extra = 1
        }

        if extra == 0 {
            break;
        }

        t := (digits[i] + extra) % 10
        extra = (digits[i] + extra) / 10
        digits[i] = t
    }

    if extra != 0 {
        digits = append([]int{extra}, digits...)
    }

    return digits
}
