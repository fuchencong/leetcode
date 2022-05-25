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
    left, right int
}

type ans61 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question61{
		question61{
			para61{
				[]int{1, 2, 3, 4, 5},
                2,4,
			},
			ans61{
				[]int{1, 4, 3, 2, 5},
			},
		},

		question61{
			para61{
				[]int{5},
                1,1,
			},
			ans61{
				[]int{5},
			},
		},

		question61{
			para61{
				[]int{1,2,3,4,5},
                1,5,
			},
			ans61{
				[]int{5,4,3,2,1},
			},
		},

		question61{
			para61{
				[]int{1, 2, 3, 4, 5},
                1,2,
			},
			ans61{
				[]int{2,1, 3, 4, 5},
			},
		},

		question61{
			para61{
				[]int{1,2},
                1,2,
			},
			ans61{
				[]int{2,1},
			},
		},
	}

	for _, c := range cases {
		l := sliceToList(c.para61.l)
		fmt.Printf("Input: %s %v %v\n", l, c.left, c.right)
		expect := sliceToList(c.ans61.result)
		fmt.Printf("Expect: %s\n", expect)

		realResult := reverseBetween(l, c.left, c.right)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
