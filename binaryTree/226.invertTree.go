package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 226. 翻转二叉树
* @Date:   2022/1/7 22:16
 */
func invertTree(root *TreeNode) *TreeNode {
	// 方法一: 递归(前序遍历)
	if root == nil {
		return nil
	}
	// 翻转
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	// 方法二: 迭代(栈进行dfs)
	st := list.New()
	if root != nil {
		st.PushBack(root)
	}
	// 中 右左，中左右
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return root
}

func invertTree3(root *TreeNode) *TreeNode {
	// 方法三: 层序遍历(队列)
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tmp := node.Left
			node.Left = node.Right
			node.Right = tmp
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}