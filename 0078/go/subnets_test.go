package leetcode

import (
	"fmt"
	"testing"
)

type question78 struct {
	para78
	ans78
}

type para78 struct {
	nums   []int
}

type ans78 struct {
	result [][]int
}

func TestPermute(t *testing.T) {
	qs := []question78{
		{
			para78{[]int{1,2,3}},
			ans78{[][]int{
                {},
                {1},
                {2},
                {3},
                {1,2},
                {1,3},
                {2,3},
                {1,2,3},
            }},
        },

		{
			para78{[]int{0}},
			ans78{[][]int{
                {},
                {0},
            }},
        },

		{
			para78{[]int{1,2}},
			ans78{[][]int{
                {},
                {1},
                {2},
                {1,2},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para78)
		fmt.Printf("Expect: %v\n", q.ans78)
		fmt.Printf("Result: %v\n\n", subsets(q.nums))
	}
}
