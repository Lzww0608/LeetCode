const N int = 26
type Trie struct {
    isEnd bool 
    children [N]*Trie
}

func (c *Trie) insert(s string) {
    cur := c 
    for i := range s {
        x := int(s[i] - 'a')
        if cur.children[x] == nil {
            cur.children[x] = &Trie{}
        }
        cur = cur.children[x]
    }
    cur.isEnd = true
    return
}

func longestWord(words []string) (ans string) {
    root := &Trie{}
    root.isEnd = true
    for _, word := range words {
        root.insert(word)
    }

    var dfs func(*Trie, []byte) 
    dfs = func(cur *Trie, s []byte) {
        if !cur.isEnd {
            return
        }
        if len(s) > len(ans) || len(s) == len(ans) && string(s) < ans {
            ans = string(s)
        }
        
        for i, v := range cur.children {
            if v != nil {
                s = append(s, byte('a' + i))
                dfs(v, s)
                s = s[:len(s)-1]
            }
        }
        return
    }
    dfs(root, []byte{})

    return
}