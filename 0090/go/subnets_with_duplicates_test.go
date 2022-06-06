package leetcode

import (
	"fmt"
	"testing"
)

type question90 struct {
	para90
	ans90
}

type para90 struct {
	nums   []int
}

type ans90 struct {
	result [][]int
}

func TestPermute(t *testing.T) {
	qs := []question90{
		{
			para90{[]int{1,2,3}},
			ans90{[][]int{
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
			para90{[]int{0}},
			ans90{[][]int{
                {},
                {0},
            }},
        },

		{
			para90{[]int{1,2}},
			ans90{[][]int{
                {},
                {1},
                {2},
                {1,2},
            }},
        },

		{
			para90{[]int{1,2,2}},
			ans90{[][]int{
                {},
                {1},
                {2},
                {1,2},
                {2,2},
                {1,2,2},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para90)
		fmt.Printf("Expect: %v\n", q.ans90)
		fmt.Printf("Result: %v\n\n", subsetsWithDup(q.nums))
	}
}
