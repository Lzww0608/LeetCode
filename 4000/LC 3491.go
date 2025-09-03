type Trie struct {
    isEnd bool 
    children [10]*Trie
}

func (c *Trie) query(s string) bool {
    cur := c 
    for i := range s {
        x := int(s[i] - '0')
        if cur.children[x] == nil {
            return true
        }
        cur = cur.children[x]
        if cur.isEnd {
            return true
        }
    }
    return false
}

func (c *Trie) insert(s string) bool {
    cur := c 
    for i := range s {
        x := int(s[i] - '0')
        if cur.children[x] == nil {
            cur.children[x] = &Trie{}
        }
        cur = cur.children[x]
    }
    if cur.isEnd {
        return false
    }
    cur.isEnd = true
    return true
}

func phonePrefix(numbers []string) bool {
    root := &Trie{}
    for _, s := range numbers {
        if !root.insert(s) {
            return false
        }
    }

    for _, s := range numbers {
        if root.query(s[:len(s)-1]) {
            return false
        }
    }

    return true
}