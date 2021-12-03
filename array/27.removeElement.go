package array

/**
* @Author: Cara1205
* @Description:给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
* @Date:   2021/9/13 22:11
 */
func removeElement(nums []int, val int) int {
	//方法：双指针
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
