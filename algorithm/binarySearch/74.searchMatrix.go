package binarySearch

/**
* @Author: Cara1205
* @Description: 力扣74.搜索二维矩阵
* @Date:   2021/10/17 20:03
 */
func searchMatrix(matrix [][]int, target int) bool {
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		//从右上角开始  i为行，j为列
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			//如果target < matrix[i][j]，则向左移（因为右边、下边的值更大，左边的值才更小，才可能找到target）
			j--
		} else {
			//如果target > matrix[i][j]，则向下移（因为下面的值更大，才能更找到target）
			i++
		}
	}
	return false
}
