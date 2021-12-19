package backtracing

/**
* @Author: Cara1205
* @Description: 力扣47. 全排列II
* 给定一个可包含重复数字的序列nums，按任意顺序返回所有不重复的全排列。
* @Date:   2021/12/19 23:43
 */
var res8 [][]int
var tmpPermute [][]int
func permuteUnique(nums []int) [][]int {
	res8 = make([][]int, 0)
	tmpPermute = make([][]int, 0)
	used := make([]bool, len(nums))
	permuteUniqueDfs([]int{}, nums, used)
	return res8
}
func permuteUniqueDfs(curInts, nums []int, used []bool) {
	for _, permute := range tmpPermute {
		if IsSliceSame(permute, curInts) {
			// 重复的选择，剪枝
			return
		}
	}
	// 保存每个排列过程中的选择，用于剪枝相同选择
	val := make([]int, len(curInts))
	copy(val, curInts)
	tmpPermute = append(tmpPermute, val)

	// 返回排列
	if len(curInts) == len(nums) {
		tmp := make([]int, len(curInts))
		copy(tmp, curInts)
		res8 = append(res8, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			curInts = append(curInts, nums[i])
			used[i] = true
			permuteUniqueDfs(curInts, nums, used)
			// 回溯 恢复
			curInts = curInts[:len(curInts)-1]
			used[i] = false
		}
	}
}
func IsSliceSame(a, b []int) bool {
	// 检查两个切片是否相等（值的顺序一样）
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}