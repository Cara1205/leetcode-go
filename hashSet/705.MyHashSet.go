package hashSet

/**
* @Author: Cara1205
* @Description: 力扣705. 设计哈希集合
* @Date:   2021/9/26 14:27
 */
type MyHashSet struct {
	Set map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Set: make(map[int]int),
	}
}

func (this *MyHashSet) Add(key int)  {
	if this.Set[key] == 0 {
		this.Set[key]++
	}
}

func (this *MyHashSet) Remove(key int)  {
	if this.Set[key] != 0 {
		this.Set[key]--
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.Set[key] == 1
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */