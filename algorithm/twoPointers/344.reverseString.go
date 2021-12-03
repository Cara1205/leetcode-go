package twoPointers

/**
* @Author: Cara1205
* @Description: 力扣344.反转字符串 字符串就地逆序
* @Date:   2021/9/26 19:40
 */
func reverseString(s []byte)  {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
