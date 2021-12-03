package binarySearch

/**
* @Author: Cara1205
* @Description: 力扣35.寻找峰值（任意一个即可）
* @Date:   2021/10/17 18:00
 */
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right - left) / 2 + left
		//比较mid与mid+1的值，因为mid-1会out of range...
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}