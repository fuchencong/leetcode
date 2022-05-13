package leetcode

import (
	"fmt"
	"testing"
)

type question34 struct {
	para34
	ans34
}

type para34 struct {
	nums   []int
    target int
}

type ans34 struct {
	result []int
}

func TestThreeSum(t *testing.T) {
	qs := []question34{
		{
			para34{
                []int{5,7,7,8,8,10},
                8,
            },

			ans34{
                []int{3, 4},
            },
        },

		{
			para34{
                []int{5,7,7,8,8,10},
                6,
            },

			ans34{
                []int{-1, -1},
            },
        },

		{
			para34{
                []int{},
                0,
            },

			ans34{
                []int{-1, -1},
            },
        },

		{
			para34{
                []int{1, 1},
                1,
            },

			ans34{
                []int{0, 1},
            },
        },

		{
			para34{
                []int{1},
                1,
            },

			ans34{
                []int{0, 0},
            },
        },

		{
			para34{
                []int{10000},
                1,
            },

			ans34{
                []int{-1, -1},
            },
        },


    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para34)
		fmt.Printf("Expect: %v\n", q.ans34)
		fmt.Printf("Result: %v\n\n", searchRange(q.nums, q.target))
	}
}
