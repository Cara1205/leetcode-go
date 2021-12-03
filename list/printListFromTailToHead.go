package list

/**
* @Author: Cara1205
* @Description: 从头到尾输出链表（不对链本身进行逆序）
* @Date:   2021/9/5 19:15
* Definition for singly-linked list.
*/

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

type StackP struct {
	nums []int
}
func ConstructorStackP() StackP {
	stackP := StackP{}
	stackP.nums = []int{}
	return stackP
}
func (this *StackP) Push(val int) {
	this.nums = append(this.nums, val)
}
func (this *StackP) Pop() int{
	res := this.nums[len(this.nums) - 1]
	this.nums = this.nums[0:len(this.nums) - 1]
	return res
}
func (this *StackP) IsEmpty() bool{
	return len(this.nums) == 0
}
func ReversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	stack := ConstructorStackP()
	p := head
	for {
		stack.Push(p.Val)	// 值入栈(倒序)
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	var result []int
	for {
		result = append(result, stack.Pop())
		if stack.IsEmpty() {
			break
		}
	}
	return result
}