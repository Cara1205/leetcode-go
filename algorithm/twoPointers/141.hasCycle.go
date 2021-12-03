package twoPointers

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/9/26 19:13
 */

type ListNode struct {
	Val int
 	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	//方法二：快慢指针
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		//防止 fast.Next==nil，引发 nil.Next 异常
		//fast 每次移动2步， slow 每次移动1步
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	//方法一：哈希表（额外空间存储遍历结点，判断是否遍历过）
	hashSet := make(map[*ListNode]int)
	p := head
	for p != nil {
		if hashSet[p] != 0 {
			return true
		}
		hashSet[p]++
		p = p.Next
	}
	return false
}
