package list

/**
* @Author: Cara1205
* @Description: 力扣2. 两数相加
* 给你两个非空的链表，表示两个非负的整数。
* 它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
* 请你将两个数相加，并以相同形式返回一个表示和的链表。
* @Date:   2021/12/21 12:52
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 创建头结点newHead
	newHead := new(ListNode)
	p := newHead        // p指向新建链表正在求和的结点

	p1, p2 := l1, l2    // p1, p2指向两条链表遍历处
	c := 0  // c是进位
	for p1 != nil && p2 != nil {
		// 创建新结点，存放值
		p.Next = new(ListNode)
		p.Next.Val = (p1.Val + p2.Val + c) % 10
		c = (p1.Val + p2.Val + c) / 10

		// 向后遍历
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	// 连接剩下链表
	for p1 != nil {
		p.Next = new(ListNode)
		p.Next.Val = (p1.Val + c) % 10
		c = (p1.Val + c) / 10
		p = p.Next
		p1 = p1.Next
	}

	for p2 != nil {
		p.Next = new(ListNode)
		p.Next.Val = (p2.Val + c) % 10
		c = (p2.Val + c) / 10
		p = p.Next
		p2 = p2.Next
	}

	if c != 0 {
		p.Next = new(ListNode)
		p.Next.Val = c
		p = p.Next
	}

	return newHead.Next
}