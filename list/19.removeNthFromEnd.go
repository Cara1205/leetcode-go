package list

/**
* @Author: Cara1205
* @Description: 力扣19. 删除链表的倒数第N个结点
* 给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
* @Date:   2021/12/21 13:14
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 思路一：遍历链表得知长度，再删除对应结点
	// 若删除头结点，则单独处理
	if head == nil {
		return head
	}

	// 遍历求链表长度
	p := head
	count := 0
	for p != nil {
		p = p.Next
		count++
	}

	// 若n等于链表长度，则删除头结点
	if count == n {
		return head.Next
	}

	// 若n不等于链表长度，则找到倒数第i+1个结点
	p = head
	for i := 1; i < count-n; i++ {
		p = p.Next
	}

	p.Next = p.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// 思路二：双指针找到倒数第i+1个结点
	if head == nil {
		return head
	}

	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}

	// 若p指向nil，则说明删除头结点
	if p == nil {
		return head.Next
	}

	// 若p不为nil，则找到倒数第i+1个结点（由pre指向）
	pre := head
	for p.Next != nil {
		pre = pre.Next
		p = p.Next
	}
	// 改链，删除倒数第i个结点
	pre.Next = pre.Next.Next
	return head
}