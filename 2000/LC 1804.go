type Trie struct {
    isEnd int
    cnt int
    children [26]*Trie
}


func Constructor() Trie {
    root := Trie{}
    for i := range root.children {
        root.children[i] = &Trie{}
    }
    return root
}


func (c *Trie) Insert(word string)  {
    node := c
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &Trie{
                isEnd: 0,
                cnt: 0,
                children: [26]*Trie{},
            }
        }
        node.cnt++
        node = node.children[idx]
    }
    node.cnt++
    node.isEnd++
}


func (c *Trie) CountWordsEqualTo(word string) int {
    node := c
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return 0
        }
        node = node.children[idx]
    }
    return node.isEnd
}


func (c *Trie) CountWordsStartingWith(prefix string) int {
    node := c
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return 0
        }
        node = node.children[idx]
    }
    return node.cnt
}


func (c *Trie) Erase(word string)  {
    node := c
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return
        }
        node.cnt--
        node = node.children[idx]
    }
    if node.isEnd > 0 {
        node.cnt--
        node.isEnd--
    }
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */