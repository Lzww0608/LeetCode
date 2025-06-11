const N int = 26
type Trie struct {
    isEnd bool
    children [N]*Trie
}

func (root *Trie) insert(s string) {
    cur := root 
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
    for _, word := range words {
        root.insert(word)
    }

    var query func(cur *Trie, s string) 
    query = func(cur *Trie, s string) {
        if len(s) > len(ans) || len(ans) == len(s) && s < ans {
            ans = s
        }
        for i := range N {
            nxt := cur.children[i]
            if nxt != nil && nxt.isEnd {
                query(nxt, s + string(byte('a' + i)))
            }
        }
    }
    query(root, "")

    return 
}