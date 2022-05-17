package leetcode

import (
	"fmt"
	"testing"
)

type question47 struct {
	para47
	ans47
}

type para47 struct {
	nums   []int
}

type ans47 struct {
	result [][]int
}

func TestPermute(t *testing.T) {
	qs := []question47{
		{
			para47{[]int{1,2,3}},
			ans47{[][]int{
                {1,2,3},
                {1,3,2},
                {2,1,3},
                {2,3,1},
                {3,1,2},
                {3,2,1},
            }},
        },

		{
			para47{[]int{}},
			ans47{[][]int{
                {},
            }},
        },

		{
			para47{[]int{1}},
			ans47{[][]int{
                {1},
            }},
        },

		{
			para47{[]int{0,1}},
			ans47{[][]int{
                {0,1},
                {1,0},
            }},
        },

		{
			para47{[]int{1,1,2}},
			ans47{[][]int{
                {1,1,2},
                {1,2,1},
                {2,1,1},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para47)
		fmt.Printf("Expect: %v\n", q.ans47)
		fmt.Printf("Result: %v\n\n", permuteUnique(q.nums))
	}
}
