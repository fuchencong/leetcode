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

type question2 struct {
	para2
	ans2
}

type para2 struct {
	l1 []int
	l2 []int
}

type ans2 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question2{
		question2{
			para2{
				[]int{2, 4, 3},
				[]int{5, 6, 4},
			},
			ans2{
				[]int{7, 0, 8},
			},
		},

		question2{
			para2{
				[]int{0},
				[]int{0},
			},
			ans2{
				[]int{0},
			},
		},

		question2{
			para2{
				[]int{9, 9, 9, 9, 9, 9, 9},
				[]int{9, 9, 9, 9},
			},
			ans2{
				[]int{8, 9, 9, 9, 0, 0, 0, 1},
			},
		},
	}

	for _, c := range cases {
		l1 := sliceToList(c.para2.l1)
		l2 := sliceToList(c.para2.l2)
		expect := sliceToList(c.ans2.result)
		realResult := addTwoNumbers(l1, l2)

		fmt.Printf("Input: %s %s\n", l1, l2)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
