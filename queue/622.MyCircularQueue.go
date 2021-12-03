package queue

/**
* @Author: Cara1205
* @Description: 循环队列
* @Date:   2021/9/18 9:40
 */
type MyCircularQueue struct {
	Size int
	Nums []int
	Head int
	Tail int
	Length int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Size: k,
		Nums: make([]int, k),
		Head: 0,
		Tail: 0,
		Length: 0,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Nums[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Size
	this.Length++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % this.Size
	this.Length--
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {return -1}
	return this.Nums[this.Head]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {return -1}
	return this.Nums[(this.Tail-1+this.Size) % this.Size]
}


func (this *MyCircularQueue) IsEmpty() bool {
	if this.Length == 0 {
		return true
	}
	return false
}


func (this *MyCircularQueue) IsFull() bool {
	if this.Length == this.Size {
		return true
	}
	return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */