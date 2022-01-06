package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 199. 二叉树的右视图
* @Date:   2022/1/6 22:58
 */
func rightSideView(root *TreeNode) []int {
	// 思路：使用层序遍历，取每层最→一个结点的值
	level := make([][]int, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 层序遍历
	queue := list.New()
	queue.PushBack(root)
	var arr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level = append(level, arr)
		arr = []int{}
	}
	// 取每层最右结点的值
	for i := 0; i < len(level); i++ {
		res = append(res, level[i][len(level[i])-1])
	}
	return res
}