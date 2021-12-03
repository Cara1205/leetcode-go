package slidingWindow

/**
* @Author: Cara1205
* @Description:
* @Date:   2021/10/17 20:23
 */
func MinSubArrayLen(target int, nums []int) int {
	sum, length := 0, 0
	minLen := len(nums)
	cover := false
	for i, j := 0, 0; i < len(nums) && j < len(nums); j++ {
		//滑动窗口 向右增宽1格
		sum += nums[j]
		length++
		for sum >= target {
			cover = true
			//更新minLen
			if length < minLen {
				minLen = length
				if minLen == 1 {
					return 1
				}
			}
			//如果sum超过target了，则把最左边的格子丢弃（用i指向）
			sum -= nums[i]
			i++
			length--
		}
	}
	if !cover {
		//如果遍历完最后一格，sum仍不超过target，说明数组总和不超过target
		return 0
	}
	return minLen
}