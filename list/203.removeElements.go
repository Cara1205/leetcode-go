package list

/**
* @Author: Cara1205
* @Description:给定一个链表的头节点 head (不带头结点)和一个整数 val ，删除链表中所有满足 Node.val == val 的节点，并返回新的头节点。
* @Date:   2021/9/14 9:24
 */

//Definition for singly-linked list.
type ListNode struct {
 Val int
 Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode   //指向遍历的前一个结点
	cur := head         //指向当前遍历到的结点
	newHead := head     //指向新的head结点
	for cur != nil {
		if cur.Val == val {     //当遍历到元素等于val时
			if cur == newHead { //如果当前结点是头结点
				newHead = cur.Next  //移除当前结点，更新头结点
				cur = newHead
			} else {            //移除当前结点
				pre.Next = cur.Next
				cur = cur.Next
			}
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return newHead
}