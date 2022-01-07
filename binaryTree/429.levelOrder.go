package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 429. N叉树的层序遍历
* @Date:   2022/1/7 22:14
 */
func nLevelOrder(root *Node) [][]int {
	// 层序遍历，区别为每个结点可能有N个孩子，用for range遍历试试
	res := make([][]int, 0)
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		length := queue.Len()
		arr := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			arr = append(arr, node.Val)
			for _, child := range node.Children {
				queue.PushBack(child)
			}
		}
		res = append(res, arr)
	}
	return res
}