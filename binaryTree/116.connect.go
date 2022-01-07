package binaryTree

import "container/list"

/**
* @Author: Cara1205
* @Description: 116. 填充每个节点的下一个右侧节点指针
* @Date:   2022/1/7 22:21
 */
type _Node struct {
	Val int
	Left *_Node
	Right *_Node
	Next *_Node
}

func connect(root *_Node) *_Node {
	// 层序遍历
	level := [][]*_Node{}
	queue := list.New()
	if root == nil {
		return root
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		arr := []*_Node{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*_Node)
			arr = append(arr, node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level = append(level, arr)
	}
	// 处理Next域
	for i := 0; i < len(level); i++ {
		for j := 0; j < len(level[i]) - 1; j++ {
			level[i][j].Next = level[i][j+1]
		}
	}
	return root
}