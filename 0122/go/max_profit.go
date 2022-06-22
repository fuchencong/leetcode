package leetcode

func maxProfit(prices []int) int {
    nums := len(prices)
    if nums <= 0 {
        return 0
    }

    results := make([][2]int, nums, nums)
    results[0][0] = 0
    results[0][1] = -prices[0]

    for i := 1; i < nums; i++ {
        results[i][0] = max(results[i-1][0], results[i-1][1]+prices[i])
        results[i][1] = max(results[i-1][0]-prices[i], results[i-1][1])
    }

    return results[nums-1][0]
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
