package trie

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

// Search /** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
