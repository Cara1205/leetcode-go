package stack

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/9/25 16:30
 */
type Stack struct {
	Nums []int
	Size int
}

func (s *Stack) Push(v int) {
	s.Nums = append(s.Nums, v)
	s.Size++
}

func (s *Stack) Pop() int {
	top := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1]
	s.Size--
	return top
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

type MyQueue struct {
	s1 Stack
	s2 Stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int)  {
	this.s1.Push(x)
}


func (this *MyQueue) Pop() int {
	if this.s1.IsEmpty() {
		return -1
	}
	for !this.s1.IsEmpty() {
		this.s2.Push(this.s1.Pop())
	}
	head := this.s2.Pop()
	for !this.s2.IsEmpty() {
		this.s1.Push(this.s2.Pop())
	}
	return head
}

func (this *MyQueue) Peek() int {
	if this.s1.IsEmpty() {
		return -1
	}
	for !this.s1.IsEmpty() {
		this.s2.Push(this.s1.Pop())
	}
	head := this.s2.Nums[len(this.s2.Nums)-1]
	for !this.s2.IsEmpty() {
		this.s1.Push(this.s2.Pop())
	}
	return head
}


func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty()
}