package leetcode

func maxProfit(prices []int) int {
    profit := 0

    currMin := 10000 + 1
    for i := 0; i < len(prices); i++ {
        if prices[i] < currMin {
            currMin = prices[i]
        } else if prices[i] - currMin > profit {
            profit = prices[i] - currMin
        }
    }

    return profit
}
