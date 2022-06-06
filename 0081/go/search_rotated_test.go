package leetcode

import (
	"fmt"
	"testing"
)

type question81 struct {
	para81
	ans81
}

type para81 struct {
	nums   []int
    target int
}

type ans81 struct {
	result bool
}

func TestThreeSum(t *testing.T) {
	qs := []question81{
		{
			para81{
                []int{2,5,6,0,0,1,2},
                0,
            },

			ans81{
                true,
            },
        },

		{
			para81{
                []int{2,5,6,0,0,1,2},
                3,
            },

			ans81{
                false,
            },
        },

		{
			para81{
                []int{},
                0,
            },

			ans81{
                false,
            },
        },

		{
			para81{
                []int{1, 1},
                1,
            },

			ans81{
                true,
            },
        },

		{
			para81{
                []int{1},
                1,
            },

			ans81{
                true,
            },
        },

		{
			para81{
                []int{2,5,6,0,0,1,2},
                2,
            },

			ans81{
                true,
            },
        },

		{
			para81{
                []int{2,5,6,0,0,1,2},
                6,
            },

			ans81{
                true,
            },
        },

		{
			para81{
                []int{2,5,6,0,0,1,2},
                7,
            },

			ans81{
                false,
            },
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para81)
		fmt.Printf("Expect: %v\n", q.ans81)
		fmt.Printf("Result: %v\n\n", search(q.nums, q.target))
	}
}
