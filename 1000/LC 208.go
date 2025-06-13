type Trie struct {
    isEnd bool 
    children [N]*Trie
}
const N int = 26

func Constructor() Trie {
    return Trie{}
}


func (c *Trie) Insert(word string)  {
    cur := c 
    for _, c := range word {
        x := int(c - 'a')
        if cur.children[x] == nil {
            cur.children[x] = &Trie{}
        }
        cur = cur.children[x]
    }
    cur.isEnd = true
}


func (c *Trie) Search(word string) bool {
    cur := c 
    for _, c := range word {
        x := int(c - 'a')
        if cur.children[x] == nil {
            return false
        }
        cur = cur.children[x]
    }
    return cur.isEnd
}


func (c *Trie) StartsWith(prefix string) bool {
    cur := c 
    for _, c := range prefix {
        x := int(c - 'a')
        if cur.children[x] == nil {
            return false
        }
        cur = cur.children[x]
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