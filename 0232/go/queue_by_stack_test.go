package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)                     // queue is: [1]
	myQueue.Push(2)                     // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Printf("%v\n", myQueue.Peek())  // return 1
	fmt.Printf("%v\n", myQueue.Pop())   // return 1, queue is [2]
	fmt.Printf("%v\n", myQueue.Empty()) // return false
}
