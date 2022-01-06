package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 107. 二叉树的层序遍历II
* @Date:   2022/1/6 22:57
 */
func levelOrderBottom(root *TreeNode) [][]int {
	// 使用队列进行层序遍历
	res := make([][]int, 0)
	que := list.New()
	if root != nil {
		que.PushBack(root)
	}
	for que.Len() > 0 {
		length := que.Len()
		var arr []int
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
	// 逆序res
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l, r = l+1, r-1
	}
	return res
}