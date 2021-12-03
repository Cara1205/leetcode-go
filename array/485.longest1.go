package array

/**
* @Author: Cara1205
* @Description:给定一个二进制数组， 计算其中最大连续 1 的个数
* @Date:   2021/9/13 10:08
 */
func findMaxConsecutiveOnes(nums []int) int {
	res := 0	//记录结果 最长连续1的个数
	tmp := 0 	//记录当前遍历到的连续1的个数，当遇到0时置0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tmp = 0
		} else {
			tmp++
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}