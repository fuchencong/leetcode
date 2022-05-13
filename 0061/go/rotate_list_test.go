package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{}
	prev := head

	for _, v := range nums {
		n := &ListNode{Val: v}
		prev.Next = n
		prev = n
	}

	return head.Next
}

func (l *ListNode) String() string {
	s := []string{}
	for l != nil {
		s = append(s, fmt.Sprintf("%v", l.Val))
		l = l.Next
	}
	return strings.Join(s, "->")
}

type question61 struct {
	para61
	ans61
}

type para61 struct {
	l []int
    k int
}

type ans61 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question61{
		question61{
			para61{
				[]int{1, 2, 3, 4, 5},
                2,
			},
			ans61{
				[]int{4, 5, 1, 2, 3},
			},
		},

		question61{
			para61{
				[]int{0, 1, 2},
                4,
			},
			ans61{
				[]int{2, 0, 1},
			},
		},

		question61{
			para61{
				[]int{1, 2, 3, 4, 5},
                1,
			},
			ans61{
				[]int{5, 1, 2, 3, 4},
			},
		},

		question61{
			para61{
				[]int{1, 2, 3, 4, 5},
                10,
			},
			ans61{
				[]int{1, 2, 3, 4, 5},
			},
		},

		question61{
			para61{
				[]int{1},
                10,
			},
			ans61{
				[]int{1},
			},
		},

		question61{
			para61{
				[]int{1, 2},
                9,
			},
			ans61{
				[]int{2, 1},
			},
		},

		question61{
			para61{
				[]int{1, 2},
                0,
			},
			ans61{
				[]int{2, 1},
			},
		},

		question61{
			para61{
				[]int{},
                10,
			},
			ans61{
				[]int{},
			},
		},


	}

	for _, c := range cases {
		l := sliceToList(c.para61.l)
		expect := sliceToList(c.ans61.result)
		realResult := rotateRight(l, c.k)

		fmt.Printf("Input: %s %v\n", l, c.k)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
