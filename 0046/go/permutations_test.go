package leetcode

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

type para46 struct {
	nums   []int
}

type ans46 struct {
	result [][]int
}

func TestPermute(t *testing.T) {
	qs := []question46{
		{
			para46{[]int{1,2,3}},
			ans46{[][]int{
                {1,2,3},
                {1,3,2},
                {2,1,3},
                {2,3,1},
                {3,1,2},
                {3,2,1},
            }},
        },

		{
			para46{[]int{}},
			ans46{[][]int{
                {},
            }},
        },

		{
			para46{[]int{1}},
			ans46{[][]int{
                {1},
            }},
        },

		{
			para46{[]int{0,1}},
			ans46{[][]int{
                {0,1},
                {1,0},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para46)
		fmt.Printf("Expect: %v\n", q.ans46)
		fmt.Printf("Result: %v\n\n", permute(q.nums))
	}
}
