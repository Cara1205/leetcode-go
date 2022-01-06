package list

/**
* @Author: Cara1205
* @Description: 力扣83. 删除排序链表中的重复元素
* 按升序排列的链表，给定链表的头节点head，删除所有重复的元素，使每个元素只出现一次
* @Date:   2022/1/4 19:43
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 方法：双指针
	// 思路：检查前后两个相邻结点的Val是否相同，相同则改链
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
