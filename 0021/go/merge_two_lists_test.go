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

type question21 struct {
	para21
	ans21
}

type para21 struct {
	l1 []int
	l2 []int
}

type ans21 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question21{
		question21{
			para21{
				[]int{1, 2, 4},
				[]int{1, 3, 4},
			},
			ans21{
				[]int{1, 1, 2, 3, 4, 4},
			},
		},

		question21{
			para21{
				[]int{},
				[]int{},
			},
			ans21{
				[]int{},
			},
		},

		question21{
			para21{
				[]int{1},
				[]int{1},
			},
			ans21{
				[]int{1, 1},
			},
		},
	}

	for _, c := range cases {
		l1 := sliceToList(c.para21.l1)
		l2 := sliceToList(c.para21.l2)
		expect := sliceToList(c.ans21.result)
		realResult := mergeTwoLists(l1, l2)

		fmt.Printf("Input: %s %s\n", l1, l2)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
