package sort

import "fmt"

/**
* @Author: Cara1205
* @Description: 排序算法的实现
* @Date:   2021/11/3 9:04
 */
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr)
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}