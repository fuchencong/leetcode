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

type question82 struct {
	para82
	ans82
}

type para82 struct {
	l []int
}

type ans82 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question82{
		question82{
			para82{
				[]int{1,2,3,3,4,4,5},
			},
			ans82{
				[]int{1,2,5},
			},
		},

		question82{
			para82{
				[]int{1, 1, 1, 2, 3},
			},
			ans82{
				[]int{2, 3},
			},
		},

		question82{
			para82{
				[]int{1, 1, 1},
			},
			ans82{
				[]int{},
			},
		},

		question82{
			para82{
				[]int{},
			},
			ans82{
				[]int{},
			},
		},

		question82{
			para82{
				[]int{1, 1, 2, 2, 3, 3},
			},
			ans82{
				[]int{},
			},
		},

		question82{
			para82{
				[]int{1, 2, 2, 3, 3},
			},
			ans82{
				[]int{1},
			},
		},

	}

	for _, c := range cases {
		l := sliceToList(c.para82.l)
		expect := sliceToList(c.ans82.result)
		realResult := deleteDuplicates(l)

		fmt.Printf("Input: %s\n", l)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
