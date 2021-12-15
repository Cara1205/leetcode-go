package backtracing

/**
* @Author: Cara1205
* @Description: 力扣90.子集II
* 给定一个整数数组nums（有重复元素），返回该数组所有可能的子集。
* @Date:   2021/12/15 20:05
 */
var res4 [][]int
func subsetsWithDup(nums []int) [][]int {
	res4 = make([][]int, 0)
	dfs([]int{}, nums, 0)
	return res4
}
func dfs(curInts, nums []int, start int) {
	//fmt.Printf("dfs(%v, %v, %d)\n", curInts, nums, start)
	// start: 剩余子集的起始位置

	// 剪枝
	for _, subset := range res4 {
		if IsSliceEqual(subset, curInts) {
			return
		}
	}
	// 返回每个结点的子集
	tmp := make([]int, len(curInts))
	copy(tmp, curInts)
	res4 = append(res4, tmp)

	// 不能在for循环里剪枝，会使剪枝后正常的循环无法继续
	for i := start; i < len(nums); i++ {
		curInts = append(curInts, nums[i])
		dfs(curInts, nums, i+1)
		curInts = curInts[:len(curInts)-1]
	}
}
func IsSliceEqual(a, b []int) bool {
	// 判断两个切片是否相等（值的顺序可以不同）
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	// 需要切片值的顺序也相同
	//for i, v := range a {
	//    if v != b[i] {
	//        return false
	//    }
	//}
	hashSet := make(map[int]int, len(a))
	for _, v := range a {
		hashSet[v]++
	}
	for _, v := range b {
		hashSet[v]--
		if hashSet[v] < 0 {
			return false
		}
	}
	for _, v := range a {
		if hashSet[v] > 0 {
			return false
		}
	}
	return true
}