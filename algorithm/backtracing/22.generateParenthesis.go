package backtracing

/**
* @Author: Cara1205
* @Description: 力扣22.括号生成
* @Date:   2021/12/3 22:49
 */

var res1 []string
func generateParenthesis(n int) []string {
	// 回溯法
	// 生成所有可能的并且有效的括号组合
	curStr := ""
	generateParenthesisDfs(curStr, 0, 0, n)
	return res1
}
// 做加法
func generateParenthesisDfs(curStr string, left, right, n int) {
	// left: 使用的左括号数
	// right:使用的右括号数
	if (left == n && right == n) {
		// 返回
		res1 = append(res1, curStr)
		return
	}
	if (left < right) {
		// 剪枝：使用的右括号数大于使用的左括号数
		return
	}
	if (left < n) {
		// 产生左分支
		generateParenthesisDfs(curStr + "(", left+1, right, n)
	}
	if (right < n) {
		// 产生右分支
		generateParenthesisDfs(curStr + ")", left, right+1, n)
	}
}