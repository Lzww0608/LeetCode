type Trie struct {
    isEnd bool
    child [26]*Trie
}

func findAllConcatenatedWordsInADict(words []string) (ans []string) {
    root := &Trie{}
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })

    var solve func(string) bool
    solve = func(s string) bool {
        cur := root
        for i := range s {
            x := int(s[i] - 'a')
            if cur.child[x] == nil {
                return false
            }
            cur = cur.child[x]
            if cur.isEnd && solve(s[i+1:]) {
                return true
            }
        }
        return cur.isEnd
    }

    for _, s := range words {
        if solve(s) {
            ans = append(ans, s)
            continue
        }
        cur := root
        for j := range s {
            x := int(s[j] - 'a')
            if cur.child[x] == nil {
                cur.child[x] = &Trie{}
            }
            cur = cur.child[x] 
        }
        cur.isEnd = true
    }

    return
}