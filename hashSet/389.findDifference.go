package hashSet

/**
* @Author: Cara1205
* @Description: 力扣389. 找不同
* @Date:   2021/9/26 14:16
 */
func findTheDifference(s string, t string) byte {
	hashTable := make(map[byte]int)
	for _, c := range s {
		hashTable[byte(c)]++
	}
	for _, c := range t {
		hashTable[byte(c)]--
	}
	for _, c := range t {
		if hashTable[byte(c)] != 0 {
			return byte(c)
		}
	}
	return byte(' ')
}