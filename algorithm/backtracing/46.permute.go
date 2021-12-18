package backtracing

/**
* @Author: Cara1205
* @Description: 力扣46. 全排列
* 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
* @Date:   2021/12/18 22:27
 */
var res7 [][]int
func permute(nums []int) [][]int {
	res7 = make([][]int, 0)
	used := make([]bool, len(nums))
	permuteDfs([]int{}, nums, used)
	return res7
}
func permuteDfs(curInts, nums []int, used []bool) {
	// used: 记录已用数字
	if len(curInts) == len(nums) {
		// 返回排列
		tmp := make([]int, len(curInts))
		copy(tmp, curInts)
		res7 = append(res7, tmp)
		return
	}

	// 不需要剪枝

	for i := 0; i < len(nums); i++ {
		// 依次从左 开始取未被选取的数字
		if !used[i] {	// 若是未被选过的数字，则选中，再继续递归
			curInts = append(curInts, nums[i])
			used[i] = true
			permuteDfs(curInts, nums, used)
			// 回溯 复原
			curInts = curInts[:len(curInts)-1]
			used[i] = false
		}
	}
}
