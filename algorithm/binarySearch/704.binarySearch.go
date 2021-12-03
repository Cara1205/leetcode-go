package binarySearch

/**
* @Author: Cara1205
* @Description: 力扣704. 二分查找
* @Date:   2021/10/5 10:36
 */
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right - left) / 2 + left
		if target == nums[mid] {
			return mid
		} else if target < nums[mid]{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
