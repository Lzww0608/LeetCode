type Trie struct {
    isEnd bool
    child [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (root *Trie) Insert(word string)  {
    cur := root
    for _, c := range word {
        x := int(c - 'a')
        if cur.child[x] == nil {
            cur.child[x] = &Trie{}
        }
        cur = cur.child[x]
    }
    cur.isEnd = true
}


/** Returns if the word is in the trie. */
func (root *Trie) Search(word string) bool {
    cur := root
    for _, c := range word {
        x := int(c - 'a')
        if cur.child[x] == nil {
            return false
        }
        cur = cur.child[x]
    }
    return cur.isEnd == true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (root *Trie) StartsWith(prefix string) bool {
    cur := root
    for _, c := range prefix {
        x := int(c - 'a')
        if cur.child[x] == nil {
            return false
        }
        cur = cur.child[x]
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