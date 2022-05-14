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

type question83 struct {
	para83
	ans83
}

type para83 struct {
	l []int
}

type ans83 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question83{
		question83{
			para83{
				[]int{1,2,3,3,4,4,5},
			},
			ans83{
				[]int{1,2,3,4,5},
			},
		},

		question83{
			para83{
				[]int{1, 1, 1, 2, 3},
			},
			ans83{
				[]int{1, 2, 3},
			},
		},

		question83{
			para83{
				[]int{1, 1, 1},
			},
			ans83{
				[]int{1},
			},
		},

		question83{
			para83{
				[]int{},
			},
			ans83{
				[]int{},
			},
		},

		question83{
			para83{
				[]int{1, 1, 2, 2, 3, 3},
			},
			ans83{
				[]int{1, 2, 3},
			},
		},

		question83{
			para83{
				[]int{1, 2, 2, 3, 3},
			},
			ans83{
				[]int{1, 2, 3},
			},
		},

	}

	for _, c := range cases {
		l := sliceToList(c.para83.l)
		expect := sliceToList(c.ans83.result)
		realResult := deleteDuplicates(l)

		fmt.Printf("Input: %s\n", l)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
