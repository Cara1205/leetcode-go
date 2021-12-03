package queue

import "fmt"

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/8/31 10:06
 */
type Stack struct {
	nums []int
}
func (s *Stack) New(){
	s.nums = []int{}
}
func (s *Stack) Push(value int) {
	s.nums = append(s.nums, value)
}
func (s *Stack) Pop() int {
	if len(s.nums) == 0 {
		return -1
	}
	top := s.nums[len(s.nums) - 1]
	s.nums = s.nums[0:len(s.nums) - 1]
	//因为切片操作是左闭右开，不会取len(s.nums)-1处的值，只取到前面
	return top
}
func (s *Stack) Print() {
	for _, v := range s.nums {
		fmt.Printf("%d\t",v)
	}
	fmt.Println()
}

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

func ConstructorCQueue() CQueue {
	cQueue := CQueue{}
	cQueue.stack1 = Stack{}
	cQueue.stack2 = Stack{}
	cQueue.stack1.New()
	cQueue.stack2.New()
	return cQueue
}

func (this *CQueue) AppendTail(value int)  {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	for { //stack1倒到stack2中，逆序
		if len(this.stack1.nums) != 0 {
			this.stack2.Push(this.stack1.Pop())
		} else {
			break
		}
	} //此时，stack2的top就是原本队列的队头
	head := this.stack2.Pop()	//取出队头
	//将stack1，stack2恢复
	for {
		if len(this.stack2.nums) != 0 {
			this.stack1.Push(this.stack2.Pop())
		} else {
			break
		}
	}
	return head
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */