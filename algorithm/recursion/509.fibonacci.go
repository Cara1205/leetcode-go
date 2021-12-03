package recursion

/**
* @Author: Cara1205
* @Description: 力扣509.斐波那契数列
* @Date:   2021/10/19 9:20
 */
func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := Fib(n-1) + Fib(n-2)
	return res
}
