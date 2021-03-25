package main

import (
	"container/list"
)

//不使用任何内建的哈希表库设计一个哈希集合（HashSet）。
//
// 实现 MyHashSet 类：
//
//
// void add(key) 向哈希集合中插入值 key 。
// bool contains(key) 返回哈希集合中是否存在这个值 key 。
// void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
//
//
//
// 示例：
//
//
//输入：
//["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove
//", "contains"]
//[[], [1], [2], [1], [3], [2], [2], [2], [2]]
//输出：
//[null, null, null, true, false, null, true, null, false]
//
//解释：
//MyHashSet myHashSet = new MyHashSet();
//myHashSet.add(1);      // set = [1]
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(1); // 返回 True
//myHashSet.contains(3); // 返回 False ，（未找到）
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(2); // 返回 True
//myHashSet.remove(2);   // set = [1]
//myHashSet.contains(2); // 返回 False ，（已移除）
//
//
//
// 提示：
//
//
// 0 <= key <= 106
// 最多调用 104 次 add、remove 和 contains 。
//
//
//
//
// 进阶：你可以不使用内建的哈希集合库解决此问题吗？
// Related Topics 设计 哈希表
// 👍 138 👎 0
const base = 769

//leetcode submit region begin(Prohibit modification and deletion)
type MyHashSet struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	index := s.hash(key)
	s.data[index].PushBack(key)
}

func (s *MyHashSet) hash(key int) int {
	return key % base
}

func (s *MyHashSet) Remove(key int) {
	index := s.hash(key)
	l := s.data[index]
	for it := l.Front(); it != nil; it = it.Next() {
		if it.Value.(int) == key {
			s.data[index].Remove(it)
		}
	}
}

/** Returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	index := s.hash(key)
	l := s.data[index]
	for it := l.Front(); it != nil; it = it.Next() {
		if it.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	set := Constructor()

	set.Add(2)
	contains := set.Contains(2)

	println(contains)
	set.Remove(2)
	c2 := set.Contains(2)

	println(c2)
}
