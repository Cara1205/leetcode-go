package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 力扣145. 二叉树的后序遍历
* @Date:   2022/1/6 20:52
 */
func postorderTraversal(root *TreeNode) []int {
	// 方法一: 递归
	res := make([]int, 0)
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		Traversal(node.Left)
		Traversal(node.Right)
		res = append(res, node.Val)
	}
	Traversal(root)
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	// 方法二: 迭代(利用栈)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 思路: 中入出，左右入，中右左出，逆序得左右中
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
	reverse(res)
	return res
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}