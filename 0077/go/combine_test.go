package leetcode

import (
	"fmt"
	"testing"
)

type question77 struct {
	para77
	ans77
}

type para77 struct {
    n int
    k int
}

type ans77 struct {
	result [][]int
}

func TestPermute(t *testing.T) {
	qs := []question77{
		{
			para77{4, 2},
			ans77{[][]int{
                {1,2},
                {1,3},
                {1,4},
                {2,3},
                {2,4},
                {3,4},
            }},
        },
		{
			para77{1, 1},
			ans77{[][]int{
                {1},
            }},
        },
		{
			para77{5, 5},
			ans77{[][]int{
                {1,2,3,4,5},
            }},
        },
		{
			para77{5, 1},
			ans77{[][]int{
                {1},
                {2},
                {3},
                {4},
                {5},
            }},
        },

    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para77)
		fmt.Printf("Expect: %v\n", q.ans77)
		fmt.Printf("Result: %v\n\n", combine(q.n, q.k))
	}
}
