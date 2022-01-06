package list

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/12/21 14:11
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 思路一：将一条链拆分为两条链，再交换先后顺序合并
	if head == nil {
		return head
	}
	p := head
	list1 := new(ListNode)
	list2 := new(ListNode)
	// 将一条链上相邻结点拆到两条链中
	l1, l2 := list1, list2
	for p != nil && p.Next != nil {
		l1.Next = new(ListNode)
		l1.Next.Val = p.Val
		l1 = l1.Next
		p = p.Next

		l2.Next = new(ListNode)
		l2.Next.Val = p.Val
		l2 = l2.Next
		p = p.Next
	}
	if p != nil {
		l1.Next = new(ListNode)
		l1.Next.Val = p.Val
	}

	// 合并两条链
	newHead := new(ListNode)
	p = newHead
	l1, l2 = list1.Next, list2.Next
	for l1 != nil && l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	return newHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	// 思路二：交换相邻结点
	if head == nil || head.Next == nil {
		return head
	}
	// 交换
	newHead := head.Next
	var pre *ListNode
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		if pre != nil {
			pre.Next = next
		}
		pre = cur
		cur = pre.Next
	}

	return newHead
}