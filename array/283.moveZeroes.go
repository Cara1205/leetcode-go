package array

/**
* @Author: Cara1205
* @Description:给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
* @Date:   2021/9/13 10:39
 */
func moveZeroes1(nums []int) {
	//方法一: 两次循环
	var i int
	var j int
	for i = 0; i < len(nums); i++ {
		//第一次循环 将非零元素全部前移，j指向非零元素将要存放的位置
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	//此时j指向0要填充的位置(因为非零元素全部移动完毕)
	for i = j ; i < len(nums); i++ {
		//第二次循环 填充0
		nums[i] = 0
	}
}

func moveZeroes(nums []int) {
	//方法二：一次遍历，借用快速排序的思路(当前最优解)
	//链接：https://leetcode-cn.com/problems/move-zeroes/solution/dong-hua-yan-shi-283yi-dong-ling-by-wang_ni_ma/
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			j++
		}
	}
}

func moveZeroes3(nums []int) {
	count := 0		//记录数组中0的个数
	q := Queue{}	//队列，记录遍历到的0的索引
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {	//当遍历到0时，记录个数，索引入队
			count++
			q.Push(i)
		} else {			//当遍历到非零元素时
			if q.IsEmpty() {//如果队列为空，则元素前移count个位置
				nums[i-count] = nums[i]
			} else {		//如果队列不为空，则元素填充到0的索引位置
				index := q.Pop()
				nums[index] = nums[i]
				q.Push(i)
			}
		}
	}
	for i := len(nums) - count; i < len(nums); i++ {	//末尾count个位置填0
		nums[i] = 0
	}
}

type Queue struct {
	nums []int
}

func (q *Queue) Push(val int) {
	q.nums = append(q.nums, val)
}

func (q *Queue) Pop() int{
	front := q.nums[0]
	q.nums = q.nums[1:]
	return front
}

func (q *Queue) IsEmpty() bool{
	return len(q.nums) == 0
}