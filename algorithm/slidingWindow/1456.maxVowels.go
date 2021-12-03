package slidingWindow

/**
* @Author: Cara1205
* @Description: 力扣1456.定长子串中元音的最大数目
* @Date:   2021/10/17 21:15
 */
func MaxVowels(s string, k int) int {
	sum, maxSum := 0, 0
	for i := 0; i < k; i++ {
		//先加载前k个字母进来，后续for循环中再滑动窗口
		if isVowel(s[i]) {
			sum++
		}
	}
	maxSum = sum
	for i := k; i < len(s); i++ {
		//滑动窗口
		if isVowel(s[i]) {
			sum++
		}
		if isVowel(s[i-k]) {
			sum--
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}