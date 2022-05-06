package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	const maxStrLen = (1000 + 1)
	rows := make([][]byte, numRows, numRows)

	reverse := false
	currRow := 0
	for i := 0; i < len(s); i++ {
		rows[currRow] = append(rows[currRow], s[i])
		if currRow == 0 {
			reverse = false
		}

		if currRow == numRows-1 {
			reverse = true
		}

		if reverse {
			currRow--
		} else {
			currRow++
		}
	}

	// convert string more effectly
	result := make([]byte, 0, maxStrLen)
	for i := 0; i < numRows; i++ {
		result = append(result, rows[i]...)
	}

	return string(result)
}
