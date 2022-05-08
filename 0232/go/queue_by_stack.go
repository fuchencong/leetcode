package leetcode

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		panic("empty stack")
	}

	res := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return res
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		panic("empty stack")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (this *MyQueue) transfer() {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	this.transfer()

	if this.s2.Empty() {
		panic("empty stack")
	}

	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	this.transfer()

	if this.s2.Empty() {
		panic("empty stack")
	}

	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}
