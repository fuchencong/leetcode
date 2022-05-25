package leetcode

func numTrees(n int) int {
    return makeTrees(1, n)
}

func makeTrees(min, max int) int {
    totalNums := 0

    if min > max {
        return 1;
    }

    for i := min; i <= max; i++ {
        // use i as root
        leftCnt := makeTrees(min, i-1)
        rightCnt := makeTrees(i+1, max)

        totalNums += leftCnt * rightCnt
    }

    return totalNums;
}
