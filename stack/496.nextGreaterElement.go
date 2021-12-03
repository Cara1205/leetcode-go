package stack

/**
* @Author: Cara1205
* @Description: 496.下一个更大的元素
* @Date:   2021/9/25 14:54
 */
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	// 方法二：单调栈
	res := []int{}
	//res := make([]int, len(nums1))
	nextGreaterNumber := make(map[int]int)

	stack := []int{}
	for i := len(nums2)-1; i >= 0; i-- {	//倒着遍历
		for len(stack) != 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]	//小的数 出栈
		}
		if len(stack) != 0 {
			nextGreaterNumber[nums2[i]] = stack[len(stack)-1]
		} else {
			nextGreaterNumber[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])		//入栈
	}
	for _, v := range nums1 {
		//res[i], _ = nextGreaterNumber[v]
		res = append(res, nextGreaterNumber[v])
	}
	return res
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	// 方法一：双重循环
	res := []int{}
	for _, v1 := range nums1 {
		flag := false		//标记 是否已遍历到相同的值
		for _, v2 := range nums2 {
			if v2 == v1 {
				flag = true
				continue
			}
			if flag {
				if v2 > v1 {
					res = append(res, v2)
					flag = false	//flag置回false，标记res中填入了比v1大的v2（否则flag会为true）
					break
				}
			}
		}
		if flag {
			res = append(res, -1)
		}
	}
	return res
}