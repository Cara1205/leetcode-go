package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 111. 二叉树的最小深度
* @Date:   2022/1/7 22:18
 */
func minDepth(root *TreeNode) int {
	// 方法一: 递归
	minDep := 100000
	if root == nil {
		return 0
	}
	var depthDfs func(depth int, node *TreeNode)
	depthDfs = func(depth int, node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth < minDep {
				minDep = depth
			}
			return
		}
		depthDfs(depth + 1, node.Left)
		depthDfs(depth + 1, node.Right)
	}
	depthDfs(1, root)
	return minDep
}

func minDepth2(root *TreeNode) int {
	// 方法二: 层序遍历
	depth := 0
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		length := queue.Len()
		depth++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return depth
}