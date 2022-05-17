package leetcode

import (
    "testing"
    "fmt"
)

type question45 struct {
    para45
    ans45
}

type para45 struct {
    nums []int
}

type ans45 struct {
    result int
}

func TestJump(t *testing.T) {
    cases := []question45 {
        question45 {
            para45 {
                nums: []int{2,3,1,1,4},
            },
            ans45 {
                result:2,
            },
        },

        question45 {
            para45 {
                nums: []int{2,3,0,1,4},
            },
            ans45 {
                result:2,
            },
        },

        question45 {
            para45 {
                nums: []int{1},
            },
            ans45 {
                result:0,
            },
        },

        question45 {
            para45 {
                nums: []int{4,3,2,1,5},
            },
            ans45 {
                result:1,
            },
        },

        question45 {
            para45 {
                nums: []int{1,1,1,1,1},
            },
            ans45 {
                result:4,
            },
        },

        question45 {
            para45 {
                nums: []int{2,1},
            },
            ans45 {
                result:1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", jump(c.nums))
	}
}
