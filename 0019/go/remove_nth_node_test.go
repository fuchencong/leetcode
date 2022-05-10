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

type question19 struct {
	para19
	ans19
}

type para19 struct {
	l []int
    n int
}

type ans19 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question19{
		question19 {
			para19{
				[]int{1, 2, 3, 4, 5},
                2,
			},
			ans19{
				[]int{1, 2, 3, 5},
			},
		},
		question19 {
			para19 {
				[]int{1, 2, 3, 4, 5},
                1,
			},
			ans19{
				[]int{1, 2, 3, 4},
			},
		},
		question19 {
			para19 {
				[]int{1, 2, 3, 4, 5},
                5,
			},
			ans19{
				[]int{2, 3, 4, 5},
			},
		},
		question19 {
			para19 {
				[]int{1},
                1,
			},
			ans19{
				[]int{},
			},
		},
	}

	for _, c := range cases {
		l := sliceToList(c.para19.l)
		expect := sliceToList(c.ans19.result)
		realResult := removeNthFromEnd(l, c.para19.n)

		fmt.Printf("Input: %s %d\n", l, c.n)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
