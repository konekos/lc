package main

type Trie struct {
	isEnd    bool
	children []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		i := c - 'a'
		if node.children[i] == nil {
			node.children[i] = &Trie{}
		}
		node = node.children[i]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		i := c - 'a'
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	if !node.isEnd {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		i := c - 'a'
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
