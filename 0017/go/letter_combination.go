package leetcode

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    digitsMap := map[string]string {
        "2":"abc",
        "3":"def",
        "4":"ghi",
        "5":"jkl",
        "6":"mno",
        "7":"pqrs",
        "8":"tuv",
        "9":"wxzy",
    }

    s := digitsMap[string(digits[0])]
    results := []string{}
    for _, b := range(s) {
        tmpResults := letterCombinations(digits[1:])
        if len(tmpResults) == 0 {
            results = append(results, string(b))
        } else {
            for i := 0; i < len(tmpResults); i++ {
                results = append(results, string(b) + tmpResults[i])
            }
        }
    }

    return results
}

