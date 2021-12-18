package backtracing

import "sort"

/**
* @Author: Cara1205
* @Description: 力扣39. 组合总和
* 给定一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
* 找出 candidates 中可以使数字和为目标数 target 的所有不同组合
* @Date:   2021/12/17 15:44
 */
var res5 [][]int
func combinationSum(candidates []int, target int) [][]int {
	res5 = make([][]int, 0)
	sort.Ints(candidates)
	combinationSumDfs([]int{}, candidates, len(candidates), target)
	return res5
}
// 从大到小
func combinationSumDfs(curInts, candidates []int, end, need int) {
	// end: 标识剩余子集结束的位置
	// need: 还需要的满足target的数
	if (need == 0) {
		// need为0，凑足了target，返回
		tmp := make([]int, len(curInts))
		copy(tmp, curInts)
		res5 = append(res5, tmp)
		return
	}
	if (candidates[0] > need) {
		// 剩余的数都大于need，剪枝
		return
	}
	for i := end; i >= 0; i-- {
		// 递归树
		if (candidates[i] <= need) {
			curInts = append(curInts, candidates[i])
			combinationSumDfs(curInts, candidates, i, need - candidates[i])
			curInts = curInts[:len(curInts)-1]
		}
	}
}