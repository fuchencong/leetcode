package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
    mapResults := make(map[string][]string)

    for _, s := range(strs) {
        b := []byte(s)
        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })

        ss := string(b)
        t := mapResults[ss]
        t = append(t, s)
        mapResults[ss] = t
    }

    results := make([][]string, 0, len(strs))
    for _, value := range(mapResults) {
        results = append(results, value)
    }

    return results
}
