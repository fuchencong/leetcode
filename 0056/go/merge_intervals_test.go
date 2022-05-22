package leetcode

import (
    "testing"
    "fmt"
)

type question56 struct {
    para56
    ans56
}

type para56 struct {
    intervals [][]int
}

type ans56 struct {
    result [][]int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question56 {
        question56 {
            para56 {
                intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            },
            ans56 {
                result: [][]int{{1, 6}, {8, 10}, {15, 18}},
            },
        },

        question56 {
            para56 {
                intervals: [][]int{{1, 4}, {4, 5}},
            },
            ans56 {
                result: [][]int{{1, 5}},
            },
        },

        question56 {
            para56 {
                intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
            },
            ans56 {
                result: [][]int{{1, 5}},
            },
        },

        question56 {
            para56 {
                intervals: [][]int{{1, 2}},
            },
            ans56 {
                result: [][]int{{1, 2}},
            },
        },

    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.intervals)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", merge(c.intervals))
	}
}
