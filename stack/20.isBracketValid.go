package stack

/**
* @Author: Cara1205
* @Description: 20.有效的括号 利用栈实现
* @Date:   2021/9/25 14:26
* @Referer: https://www.cnblogs.com/gyyyl/p/13645180.html
 */
func isValid(s string) bool {
	brackets := []byte{}		//切片模拟栈
	for _, v := range s {
		switch v {
		case '(', '[', '{':		//如果是左括号，则入栈
			brackets = append(brackets, byte(v))
		case ')', ']', '}':
			if len(brackets) == 0 {
				return false	//如果当前字符为右括号，而左括号栈中为空，则返回false
			}
			if v == ')' && brackets[len(brackets)-1] == '(' ||
				v == ']' && brackets[len(brackets)-1] == '[' ||
				v == '}' && brackets[len(brackets)-1] == '{' {
				brackets = brackets[:len(brackets)-1]	//当前右括号与栈顶左括号匹配，左括号出栈
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(brackets) != 0 {
		return false
	}
	return true
}
