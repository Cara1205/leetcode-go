package list

/**
* @Author: Cara1205
* @Description: 力扣61. 旋转链表
* 给你一个链表的头节点head，旋转链表，将链表每个节点向右移动k个位置。
* @Date:   2021/12/20 23:52
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 思路：每个节点右移k个位置，等价于找到末尾k个结点，修改链表
	// k有可能大于len(list)，其有效移动位相当于k%len(list)

	// 遍历链表，得到链表长度
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}

	// 双指针找到末尾k个结点的前一个结点
	newHead := head
	front, tail := head, head
	for i := 0; i < k % length; i++ {
		tail = tail.Next
	}

	for tail.Next != nil {
		// 遍历后，front指向倒数第k+1个结点，tail指向最后一个结点
		front = front.Next
		tail = tail.Next
	}

	if front.Next != nil && tail != nil {
		newHead = front.Next
		// 改链
		front.Next = nil
		tail.Next = head
	}

	return newHead
}