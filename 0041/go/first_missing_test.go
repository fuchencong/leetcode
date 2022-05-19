package leetcode

import (
    "fmt"
    "testing"
)

type question41 struct {
    para41
    ans41
}

type para41 struct {
    nums []int
}

type ans41 struct {
    result int
}

func TestFirstMissingPositive(t *testing.T) {
    cases := []question41 {
        question41 {
            para41 {
                []int{1,2,0},
            },
            ans41 {
                3,
            },
        },
        question41 {
            para41 {
                []int{3,4,-1,1},
            },
            ans41 {
                2,
            },
        },
        question41 {
            para41 {
                []int{7,8,9,11,12},
            },
            ans41 {
                1,
            },
        },
        question41 {
            para41 {
                []int{1,2,3,4,5},
            },
            ans41 {
                6,
            },
        },
        question41 {
            para41 {
                []int{1,1,3,3,5},
            },
            ans41 {
                2,
            },
        },
        question41 {
            para41 {
                []int{5},
            },
            ans41 {
                1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para41)
		fmt.Printf("Expect: %v\n", c.ans41)
		fmt.Printf("Result: %v\n\n", firstMissingPositive(c.nums))
	}
}



