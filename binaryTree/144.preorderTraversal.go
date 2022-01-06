package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 力扣144. 二叉树的前序遍历
* @Date:   2022/1/6 19:46
 */

type TreeNode struct {
	Val		int
	Left	*TreeNode
	Right	*TreeNode
}

func preorderTraversal1(root *TreeNode) []int {
	// 方法一: 递归
	res := make([]int, 0)
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		Traversal(node.Left)
		Traversal(node.Right)
	}
	Traversal(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	// 方法二: 迭代(使用栈)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 思路：中入出,右左入栈，中左右出栈
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	return res
}