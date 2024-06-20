type Trie struct {
    child [26]*Trie
}

func minimumLengthEncoding(words []string) (ans int) {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j])
    })

    root := &Trie{}
    for _, s := range words {
        cur := root
        f := false
        for i := len(s) - 1; i >= 0; i-- {
            x := int(s[i] - 'a')
            if cur.child[x] == nil {
                f = true
                cur.child[x] = &Trie{}
            }
            cur = cur.child[x]
        }
        if f {
            ans += len(s) + 1
        }
    }

    return 
}