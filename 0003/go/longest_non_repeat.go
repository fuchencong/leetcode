package leetcode

func lengthOfLongestSubstring(s string) int {
	result, start, i := 0, 0, 0
	occurs := make(map[byte]int, len(s))

	for i < len(s) {
		pos, ok := occurs[s[i]]
		// repeat character
		if ok && pos >= start {
			start = pos + 1
		}
		result = max(result, i-start+1)
		occurs[s[i]] = i
		i++
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
