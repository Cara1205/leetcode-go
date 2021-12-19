package backtracing

/**
* @Author: Cara1205
* @Description: 剑指 Offer 38. 字符串的排列
* 输入一个字符串，打印出该字符串中字符的所有排列。
* @Date:   2021/12/19 23:45
 */
var res []string
// var tmpStr []string
var tmpStr map[string]int
func permutation(s string) []string {
	res = []string{}
	// tmpStr = []string{}
	tmpStr = make(map[string]int, 0)
	used := make([]bool, len(s))
	permutationDfs("", s, used)
	return res
}
func permutationDfs(curStr, s string, used []bool) {
	// used: 标记已使用的字符

	// 超时的去重：
	// for _, str := range tmpStr {
	//     if strings.EqualFold(str, curStr) {
	//         // 重复的字符，剪枝
	//         return
	//     }
	// }

	// 修改的去重：
	if tmpStr[curStr] > 0 {
		return
	}

	// 保存每个排列过程中的选择，用于剪枝相同选择
	// tmpStr = append(tmpStr, curStr)
	tmpStr[curStr]++

	if len(curStr) == len(s) {
		// 返回字符的排列
		res = append(res, curStr)
	}

	for i := 0; i < len(s); i++ {
		if !used[i] {
			used[i] = true
			permutationDfs(curStr + string(s[i]), s, used)
			// 回溯 恢复
			used[i] = false
		}
	}
}