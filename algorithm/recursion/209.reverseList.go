package recursion

/**
* @Author: Cara1205
* @Description: 力扣209.反转链表
* @Date:   2021/10/19 9:24
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList1(head *ListNode) *ListNode {
	//方法一：就地逆序(不带头结点的链表)
	if head == nil || head.Next == nil{
		return head
	}
	var pre *ListNode
	cur := head
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func ReverseList2(head *ListNode) *ListNode {
	//方法二：头插法(带头结点比较好做，自行添加一个头结点)
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{0, head}
	newHead.Next = head
	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = newHead.Next
		newHead.Next = cur
		cur = next
	}
	return newHead.Next
}

func ReverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}