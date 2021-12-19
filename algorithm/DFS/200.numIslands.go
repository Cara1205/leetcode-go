package DFS

/**
* @Author: Cara1205
* @Description: 力扣200. 岛屿数量
* 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
* @Date:   2021/12/19 23:47
 */

// 存在问题：本地IDE可正常运行，力扣上执行总为0
var res1 int
func numIslands(grid [][]byte) int {
	res1 = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byte(1) {
				dfs(grid, i, j)
				res1 = res1 + 1
			}
		}
	}
	return res1
}
func dfs(grid [][]byte, row, column int) {
	// row: 指向当前行
	// column: 指向当前列
	if (row < 0) || (row >= len(grid)) || (column < 0) || (column >= len(grid[0])) {
		// 判断格子是否在网格中，若超出边界范围则返回
		return
	}

	if (grid[row][column] != byte(1)) {
		// 如果该格子不是岛屿，直接返回
		return
	}
	// 是格子的，则将值置为2，表示已遍历过
	grid[row][column] = byte(2)

	// 遍历上、下、左、右四个方向的格子
	dfs(grid, row-1, column)
	dfs(grid, row+1, column)
	dfs(grid, row, column-1)
	dfs(grid, row, column+1)
}