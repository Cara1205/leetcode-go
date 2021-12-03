package binarySearch

/**
* @Author: Cara1205
* @Description: 力扣35. 搜索插入位置
* @Date:   2021/10/17 17:47
 */
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target <= nums[left] {
		return 0
	} else if target > nums[right] {
		return len(nums)
	}
	for left <= right {
		mid := (right - left) / 2 + left
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	return left
}
