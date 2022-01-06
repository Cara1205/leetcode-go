package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 102. 二叉树的层序遍历
* @Date:   2022/1/6 22:56
 */
func levelOrder(root *TreeNode) [][]int {
	// 使用队列进行层序遍历
	res := make([][]int, 0)
	que := list.New()
	if root != nil {
		que.PushBack(root)
	}
	var arr []int   // 保存本层结点切片
	for que.Len() > 0 {
		length := que.Len()     // 获取当前队列元素个数，即树的本层的结点个数
		// arr := make([]int, length)  // 为什么在这儿这么写不行?
		for i := 0; i < length; i++ {
			node := que.Remove(que.Front()).(*TreeNode)
			arr = append(arr, node.Val)
			if node.Left != nil {
				que.PushBack(node.Left)
			}
			if node.Right != nil {
				que.PushBack(node.Right)
			}

		}
		res = append(res, arr)
		arr = []int{}
	}
	return res
}