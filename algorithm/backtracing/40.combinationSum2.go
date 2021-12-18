package backtracing

import "sort"

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/12/17 16:39
 */
var res6 [][]int
var tmpSubsets [][]int
func combinationSum2(candidates []int, target int) [][]int {
	res6 = make([][]int, 0)
	tmpSubsets = make([][]int, 0)
	sort.Ints(candidates)
	combinationSum2Dfs([]int{}, candidates, len(candidates)-1, target)
	return res6
}
func combinationSum2Dfs(curInts, candidates []int, end, need int) {
	// end: 标识剩余子集的结束位置
	// need: 凑齐target还需要的数值

	// 相同组合，剪枝
	for _, subset := range tmpSubsets {
		if isSliceEqual(subset, curInts) {
			return
		}
	}

	tmp := make([]int, len(curInts))
	copy(tmp, curInts)
	tmpSubsets = append(tmpSubsets, tmp)

	if (need == 0) {
		// need为0，凑齐target，返回
		tmp := make([]int, len(curInts))
		copy(tmp, curInts)
		res6 = append(res6, tmp)
		return
	}
	if (candidates[0] > need) {
		// 剩余子集中所有数均大于need，剪枝
		return
	}
	for i := end; i >= 0; i-- {
		// 递归树
		if (candidates[i] <= need) {
			curInts = append(curInts, candidates[i])
			combinationSum2Dfs(curInts, candidates, i-1, need - candidates[i])
			curInts = curInts[:len(curInts)-1]

		}
	}
}
func isSliceEqual(a, b []int) bool {
	// 判断两个切片是否相等（值的顺序可以不同）
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	hashSet := make(map[int]int, len(a))
	for _, v := range a {
		hashSet[v]++
	}
	for _, v := range b {
		hashSet[v]--
		if hashSet[v] < 0 {
			return false
		}
	}
	for _, v := range a {
		if hashSet[v] > 0 {
			return false
		}
	}
	return true
}