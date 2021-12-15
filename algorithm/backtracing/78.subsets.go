package backtracing

/**
* @Author: Cara1205
* @Description: 力扣78.子集
* 给定一个整数数组nums（无重复元素），返回该数组所有可能的子集（幂集）。
* @Date:   2021/12/15 19:22
 */

var res2 [][]int
func subsets(nums []int) [][]int {
	res2 = make([][]int, 0)
	subsetDfs([]int{}, nums, 0)
	return res2
}
func subsetDfs(curInts, nums []int, start int) {
	// start: 剩余子集的起始位置
	// 返回每个结点的子集
	tmp := make([]int, len(curInts))
	copy(tmp, curInts)
	res2 = append(res2, tmp)

	// 不需要剪枝

	for i := start; i < len(nums); i++ {
		curInts = append(curInts, nums[i])
		subsetDfs(curInts, nums, i+1)
		curInts = curInts[:len(curInts)-1]
	}
}