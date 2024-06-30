type Trie struct {
    cnt int
    child [26]*Trie
}

type WordsFrequency struct {
    root Trie
}


func Constructor(book []string) WordsFrequency {
    root := &Trie{}
    for _, s := range book {
        cur := root
        for _, c := range s {
            x := int(c - 'a')
            if cur.child[x] == nil {
                cur.child[x] = &Trie{}
            }
            cur = cur.child[x]
        }
        cur.cnt++
    }

    return WordsFrequency{*root}
}


func (c *WordsFrequency) Get(word string) int {
    cur := &c.root
    for _, c := range word {
        x := int(c - 'a')
        if cur.child[x] == nil {
            return 0
        }
        cur = cur.child[x]
    }

    return cur.cnt
}


/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */