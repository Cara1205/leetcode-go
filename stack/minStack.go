package leetcode

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/8/31 22:06
 */
type MinStack struct {
	nums []int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	minStack := MinStack{}
	minStack.nums = []int{}
	return minStack
}

func (this *MinStack) Push(x int)  {
	this.nums = append(this.nums, x)
}


func (this *MinStack) Pop()  {
	this.nums = this.nums[0:len(this.nums)-1]
}


func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}


func (this *MinStack) Min() int {
	minnest := 1 << 31 - 1
	for _, v := range this.nums {
		if v < minnest {
			minnest = v
		}
	}
	return minnest
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */