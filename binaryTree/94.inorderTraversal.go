package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 力扣94. 二叉树的中序遍历
* @Date:   2022/1/6 20:42
 */
func inorderTraversal(root *TreeNode) []int {
	// 方法一: 递归
	res := make([]int, 0)
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		Traversal(node.Left)
		res = append(res, node.Val)
		Traversal(node.Right)
	}
	Traversal(root)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	// 方法二: 迭代(使用栈)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	st := list.New()
	cur := root
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}