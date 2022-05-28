package leetcode

import (
    "fmt"
    "testing"
)

type question66 struct {
    para66
    ans66
}

type para66 struct {
    para []int
}

type ans66 struct {
    result []int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question66 {
        question66 {
            para66 {
                []int{1, 3, 5},
            },
            ans66 {
                []int{1, 3, 6},
            },
        },
        question66 {
            para66 {
                []int{2,2,2,0,1},
            },
            ans66 {
                []int{2,2,2,0,2},
            },
        },
        question66 {
            para66 {
                []int{5,5,5,5,5},
            },
            ans66 {
                []int{5,5,5,5,6},
            },
        },
        question66 {
            para66 {
                []int{0},
            },
            ans66 {
                []int{1},
            },
        },
        question66 {
            para66 {
                []int{9},
            },
            ans66 {
                []int{1,0},
            },
        },
        question66 {
            para66 {
                []int{9,9},
            },
            ans66 {
                []int{1,0,0},
            },
        },
        question66 {
            para66 {
                []int{8,9},
            },
            ans66 {
                []int{9,0},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", plusOne(c.para))
	}
}
