package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 515. 在每个树行中找最大值
* @Date:   2022/1/7 22:22
 */
func largestValues(root *TreeNode) []int {
	// 层序遍历
	level := make([][]int, 0)
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		length := queue.Len()
		arr := []int{}
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
	}
	// 找出每层的最大值
	res := make([]int, 0)
	for i := 0; i < len(level); i++ {
		max := -2147483648
		for _, v := range level[i] {
			if v > max {
				max = v
			}
		}
		res = append(res, max)
	}
	return res
}