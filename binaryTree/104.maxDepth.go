package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 104. 二叉树的最大深度
* @Date:   2022/1/7 22:15
 */
func maxDepth(root *TreeNode) int {
	// 方法一: 递归
	maxDep := 0
	var depthDFS func(depth int, node *TreeNode)
	depthDFS = func(depth int, node *TreeNode) {
		if node == nil {
			if depth-1 > maxDep {
				maxDep = depth - 1
			}
			return
		}
		depthDFS(depth + 1, node.Left)
		depthDFS(depth + 1, node.Right)
	}
	depthDFS(1, root)
	return maxDep
}

func maxDepth2(root *TreeNode) int {
	// 方法二: 层序遍历
	level := [][]int{}
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
	return len(level)
}