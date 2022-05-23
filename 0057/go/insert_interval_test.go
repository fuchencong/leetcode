package leetcode

import (
    "testing"
    "fmt"
)

type question57 struct {
    para57
    ans57
}

type para57 struct {
    intervals [][]int
    newInterval []int
}

type ans57 struct {
    result [][]int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question57 {
        question57 {
            para57 {
                intervals: [][]int{{1, 3}, {6, 9}},
                newInterval: []int{2, 5},
            },
            ans57 {
                result: [][]int{{1, 5}, {6, 9}},
            },
        },

        question57 {
            para57 {
                intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
                newInterval: []int{4, 8},
            },
            ans57 {
                result: [][]int{{1, 2}, {3, 10}, {12, 16}},
            },
        },

        question57 {
            para57 {
                intervals: [][]int{{1, 2}, {6, 9}},
                newInterval: []int{10, 20},
            },
            ans57 {
                result: [][]int{{1, 2}, {6, 9}, {10, 20}},
            },
        },

        question57 {
            para57 {
                intervals: [][]int{{1, 2}, {6, 9}},
                newInterval: []int{1, 9},
            },
            ans57 {
                result: [][]int{{1, 9}},
            },
        },

        question57 {
            para57 {
                intervals: [][]int{{1, 2}, {6, 9}},
                newInterval: []int{1, 3},
            },
            ans57 {
                result: [][]int{{1, 3}, {6, 9}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v %v\n", c.intervals, c.newInterval)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", insert(c.intervals, c.newInterval))
	}
}
