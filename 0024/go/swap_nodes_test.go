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

type question24 struct {
	para24
	ans24
}

type para24 struct {
	l []int
}

type ans24 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question24{
		question24{
			para24{
				[]int{1, 2, 3, 4},
			},
			ans24{
				[]int{2, 1, 4, 3},
			},
		},

		question24{
			para24{
				[]int{},
			},
			ans24{
				[]int{},
			},
		},

		question24{
			para24{
				[]int{1},
			},
			ans24{
				[]int{1},
			},
		},

		question24{
			para24{
				[]int{1, 2, 3},
			},
			ans24{
				[]int{2, 1, 3},
			},
		},
	}

	for _, c := range cases {
		l := sliceToList(c.para24.l)
		fmt.Printf("Input: %s\n", l)

		expect := sliceToList(c.ans24.result)
		realResult := swapPairs(l)

		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
