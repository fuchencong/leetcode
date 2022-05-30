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

func inorder(root *TreeNode) []int {
    result := []int{}

    if root == nil {
        return result
    }

    result = append(result, inorder(root.Left)...)
    result = append(result, root.Val)
    result = append(result, inorder(root.Right)...)

    return result
}


type question109 struct {
	para109
	ans109
}

type para109 struct {
	l []int
}

type ans109 struct {
	result []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []question109{
		question109{
			para109{
				[]int{1, 2, 3, 4, 5},
			},
			ans109{
				[]int{1,2,3,4,5},
			},
		},

		question109{
			para109{
				[]int{-10,-3,0,5,9},
			},
			ans109{
				[]int{-10,-3,0,5,9},
			},
		},
		question109{
			para109{
				[]int{1},
			},
			ans109{
				[]int{1},
			},
		},
	}

	for _, c := range cases {
		l := sliceToList(c.para109.l)
		expect := sliceToList(c.ans109.result)
		realResult := sortedListToBST(l)

		fmt.Printf("Input: %s\n", l)
		fmt.Printf("Expect: %s\n", expect)
		fmt.Printf("Result: %v\n\n", inorder(realResult))
	}

}
