package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    results := make([][]int, 0, len(intervals))
    for i := 0; i < len(intervals); i++ {
        if i == 0 {
            results = append(results, intervals[i])
            continue
        }

        end := results[len(results)-1]
        if end[1] < intervals[i][0] {
            results = append(results, append([]int{}, intervals[i]...))
        } else {
            end[1] = max(end[1], intervals[i][1])
        }
    }

    return results
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
