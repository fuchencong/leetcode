package leetcode

import (
    "testing"
    "fmt"
)

type question55 struct {
    para55
    ans55
}

type para55 struct {
    nums []int
}

type ans55 struct {
    result bool
}

func TestJump(t *testing.T) {
    cases := []question55 {
        question55 {
            para55 {
                nums: []int{2,3,1,1,4},
            },
            ans55 {
                result:true,
            },
        },

        question55 {
            para55 {
                nums: []int{3,2,1,0,4},
            },
            ans55 {
                result:false,
            },
        },

        question55 {
            para55 {
                nums: []int{1},
            },
            ans55 {
                result:true,
            },
        },

        question55 {
            para55 {
                nums: []int{4,3,2,1,5},
            },
            ans55 {
                result:true,
            },
        },

        question55 {
            para55 {
                nums: []int{1,1,1,1,1},
            },
            ans55 {
                result:true,
            },
        },

        question55 {
            para55 {
                nums: []int{2,1},
            },
            ans55 {
                result:true,
            },
        },

        question55 {
            para55 {
                nums: []int{2,1,0,4},
            },
            ans55 {
                result:false,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", canJump(c.nums))
	}
}
