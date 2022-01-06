package list

/**
* @Author: Cara1205
* @Description: 力扣82.删除排序链表中的重复元素II
* 按升序排列的链表，给你这个链表的头节点head，删除链表中所有存在数字重复情况的节点，只保留原始链表中没有重复出现的数字。
* @Date:   2022/1/4 19:37
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 申请一个头结点，第一个元素重复不用单独处理
	newHead := new(ListNode)
	newHead.Next = head
	pre := newHead
	cur := head
	for cur != nil && cur.Next != nil {
		// 当元素不重复时，同时更新pre和cur
		if cur.Val != cur.Next.Val {
			pre = cur
			cur = cur.Next
		} else {
			// 当元素重复时，移动cur到最后一个重复元素处
			for cur.Next != nil && cur.Val == cur.Next.Val{
				cur = cur.Next
			}
			// pre指向重复元素的后面一个元素
			cur = cur.Next
			pre.Next = cur
		}
	}
	return newHead.Next
}