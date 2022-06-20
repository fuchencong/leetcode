package leetcode

import (
	"fmt"
	"testing"
)

type question118 struct {
	para118
	ans118
}

type para118 struct {
	nums   int
}

type ans118 struct {
	result [][]int
}

func TestThreeSum(t *testing.T) {
	qs := []question118{
		{
			para118{1},
			ans118{[][]int{
                {1},
            }},
        },
		{
			para118{5},
			ans118{[][]int{
                {1},
                {1,1},
                {1,2,1},
                {1,3,3,1},
                {1,4,6,4,1},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para118)
		fmt.Printf("Expect: %v\n", q.ans118)
		fmt.Printf("Result: %v\n\n", generate(q.nums))
	}
}
