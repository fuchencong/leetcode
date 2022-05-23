package leetcode


func insert(intervals [][]int, newInterval []int) [][]int {
    insertPos := binarySearch(intervals, newInterval)
    intervals = append(intervals[:insertPos], append([][]int{newInterval}, intervals[insertPos:]...)...)

    startPos := insertPos
    if startPos > 0 && intervals[startPos-1][1] >= intervals[startPos][0] {
        startPos--
    }

    results := mergeInterval(intervals, startPos)
    return results
}


func binarySearch(intervals [][]int, newInterval []int) int {
    left := 0;
    right := len(intervals) - 1

    for left <= right {
        mid := (left+right) / 2

        if intervals[mid][0] < newInterval[0] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left
}

func mergeInterval(intervals [][]int, startPos int) [][]int {
    results := make([][]int, 0, len(intervals))
    results = append(results, intervals[0:startPos+1]...)

    i := startPos+1
    for ; i < len(intervals); i++ {
        if intervals[i][0] <= results[startPos][1] {
            results[startPos][1] = max(intervals[i][1], results[startPos][1])
        } else {
            break
        }
    }

    results = append(results, intervals[i:]...)
    return results
}


func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
