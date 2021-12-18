package list

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/12/15 20:29
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	//方法一：原地逆序
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

func reverseList2(head *ListNode) *ListNode {
	//方法二：头插法
	newHead := &ListNode{}  //新建一个头结点，在头结点后插入被逆序的结点
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		newHead.Next = cur      //头插
		cur.Next = pre          //逆序
		pre = cur
		cur = next
	}
	return newHead.Next
}

func reverseList3(head *ListNode) *ListNode {
	//方法三：递归，将当前head连接到已逆序的子链表后
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	//key point: head.Next现在是子链表中的最后一个结点
	head.Next.Next = head
	head.Next = nil
	return newHead
}