package binaryTree

/**
* @Author: Cara1205
* @Description: 559. N叉树的最大深度
* @Date:   2022/1/7 21:41
 */
type Node struct {
	Val int
	Children []*Node
}

func nMaxDepth(root *Node) int {
	// 方法一: 递归
	maxDep := 0
	if root == nil {
		return maxDep
	}
	var depthDFS func(depth int, node *Node)
	depthDFS = func(depth int, node *Node) {
		if node.Children == nil || len(node.Children) == 0 {
			if depth > maxDep {
				maxDep = depth
			}
			return
		}
		for _, child := range node.Children {
			depthDFS(depth + 1, child)
		}
	}
	depthDFS(1, root)
	return maxDep
}