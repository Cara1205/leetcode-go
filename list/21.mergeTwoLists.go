package list

/**
* @Author: Cara1205
* @Description: 力扣21. 合并两个有序链表
* 将两个升序链表合并为一个新的升序链表并返回。
* @Date:   2021/12/20 23:31
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 头结点
	newHead := &ListNode{}
	p := newHead
	node1, node2 := list1, list2
	for ; node1 != nil && node2 != nil; {
		if node1.Val < node2.Val {
			p.Next = node1
			p = p.Next
			node1 = node1.Next
		} else {
			p.Next = node2
			p = p.Next
			node2 = node2.Next
		}
	}
	// 剩下的结点加入链表
	for node1 != nil {
		p.Next = node1
		p = p.Next
		node1 = node1.Next
	}

	for node2 != nil {
		p.Next = node2
		p = p.Next
		node2 = node2.Next
	}
	return newHead.Next
}