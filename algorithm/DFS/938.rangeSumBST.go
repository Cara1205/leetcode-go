package DFS

/**
* @Author: Cara1205
* @Description: 力扣938. 二叉搜索树的范围和
* @Date:   2021/12/19 23:46
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var res int
func rangeSumBST(root *TreeNode, low int, high int) int {
	Dfs(root, low, high)
	return res
}
func Dfs(root *TreeNode, low, high int) {
	if root == nil {
		return
	}
	if (root.Val >= low) && (root.Val <= high) {
		res = res + root.Val
	}
	Dfs(root.Left, low, high)
	Dfs(root.Right, low, high)
}