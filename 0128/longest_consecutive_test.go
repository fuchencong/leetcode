package leetcode

import (
    "testing"
    "fmt"
)

type question128 struct {
    para128
    ans128
}

type para128 struct {
    nums []int
}

type ans128 struct {
    result int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question128 {
        question128 {
            para128 {
                nums: []int{100,4,200,1,3,2},
            },
            ans128 {
                result:4,
            },
        },

        question128 {
            para128 {
                nums: []int{1},
            },
            ans128 {
                result:1,
            },
        },

        question128 {
            para128 {
                nums: []int{0,3,7,2,5,8,4,6,0,1},
            },
            ans128 {
                9,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", longestConsecutive(c.nums))
	}
}
