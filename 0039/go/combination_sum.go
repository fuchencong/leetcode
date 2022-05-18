package leetcode

func combinationSum(candidates []int, target int) [][]int {
    results := [][]int{}

    choosed := []int{}
    var dfsSearch func  (candidates []int, target int, currSum int, startPos int)
    dfsSearch = func(candidates []int, target int, currSum int, startPos int) {
        if currSum == target {
            r := make([]int, 0, len(choosed))
            results = append(results, append(r, choosed...))
            return
        }

        if currSum > target {
            return
        }

        for i := startPos; i < len(candidates); i++ {
            choosed = append(choosed, candidates[i])
            currSum += candidates[i]
            dfsSearch(candidates, target, currSum, i)
            choosed = choosed[:len(choosed)-1]
            currSum -= candidates[i]
        }
    }

    dfsSearch(candidates, target, 0, 0)
    return results
}


