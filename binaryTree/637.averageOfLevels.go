package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 637. 二叉树的层平均值
* @Date:   2022/1/6 22:59
 */
func averageOfLevels(root *TreeNode) []float64 {
	// 层序遍历，然后计算每层平均值
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	// 层序遍历
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		// arr := make([]int, length)
		sum := 0
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		// 计算每层平均值
		res = append(res, float64(sum)/float64(length))
	}
	return res
}