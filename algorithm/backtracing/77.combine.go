package backtracing

/**
* @Author: Cara1205
* @Description: 力扣77.组合
* 给定两个整数 n 和 k，返回范围[1, n]中所有可能的 k 个数的组合。
* @Date:   2021/12/15 19:53
 */

var res3 [][]int
func combine(n int, k int) [][]int {
	res3 = make([][]int, 0)
	combineDfs([]int{}, 0, 1, n, k)
	return res3
}
func combineDfs(curInts []int, num, start, n, k int) {
	// num: 已选数的个数
	// start: 剩余子集的起始位置
	if (num == k) {
		// 返回
		tmp := make([]int, len(curInts))
		copy(tmp, curInts)
		res3 = append(res3, tmp)
		return
	}

	if (n-start+1 < k-num) {
		// 剪枝：剩余子集中的数不足
		return
	}

	for i := start; i <= n; i++ {
		curInts = append(curInts, i)
		combineDfs(curInts, num+1, i+1, n, k)
		curInts = curInts[:len(curInts)-1]
	}
}