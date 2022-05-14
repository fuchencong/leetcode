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

type question86 struct {
	para86
	ans86
}

type para86 struct {
	l []int
    k int
}

type ans86 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question86{
		question86{
			para86{
				[]int{1,4,3,2,5,2},
                3,
			},
			ans86{
				[]int{1,2,2,4,3,5},
			},
		},

		question86{
			para86{
				[]int{2,1},
                2,
			},
			ans86{
				[]int{1,2},
			},
		},

		question86{
			para86{
				[]int{1},
                10,
			},
			ans86{
				[]int{1},
			},
		},

		question86{
			para86{
				[]int{1, 2, 3, 4, 5},
                3,
			},
			ans86{
				[]int{1, 2, 3, 4, 5},
			},
		},

		question86{
			para86{
				[]int{2,2,1,1,1},
                3,
			},
			ans86{
				[]int{2,2,1,1,1},
			},
		},

		question86{
			para86{
				[]int{2,2,1,1,1},
                1,
			},
			ans86{
				[]int{2,2,1,1,1},
			},
		},

		question86{
			para86{
				[]int{2,3,1,1,1},
                2,
			},
			ans86{
				[]int{1,1,1,2,3},
			},
		},

		question86{
			para86{
				[]int{},
                10,
			},
			ans86{
				[]int{},
			},
		},


	}

	for _, c := range cases {
		l := sliceToList(c.para86.l)
		fmt.Printf("Input: %s %v\n", l, c.k)
		expect := sliceToList(c.ans86.result)
		fmt.Printf("Expect: %s\n", expect)
		realResult := partition(l, c.k)
		fmt.Printf("Result: %s\n\n", realResult)
	}

}
