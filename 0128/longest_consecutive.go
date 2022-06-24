package leetcode

func longestConsecutive(nums []int) int {
    existed := make(map[int]bool)
    for _, n := range nums {
        existed[n] = true
    }

    result := 0
    for _, n := range nums {
        // if n-1 existed, we can always check n-1 to get maxium
        if existed[n-1] {
            continue
        }

        j := n+1
        for true {
            if !existed[j] {
                break
            }
            j++
        }


        result = max(j-n, result)
    }

    return result
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
