func replaceWords(dictionary []string, sentence string) string {
    type Trie struct {
        isEnd bool
        next [26]*Trie
    }
    root := &Trie{}
    for _, s := range dictionary {
        cur := root
        for _, c := range s {
            x := int(c - 'a')
            if cur.next[x] == nil {
                cur.next[x] = &Trie{}
            } 
            cur = cur.next[x]
        }
        cur.isEnd = true
    }

    builder := strings.Builder{}
    s := strings.Split(sentence, " ")
    for i, t := range s {
        j, cur := 0, root
        for j < len(t) {
            x := int(t[j] - 'a')
            if cur.next[x] == nil {
                j = len(t)
                break 
            } 
            cur = cur.next[x]
            j++
            if cur.isEnd {
                break
            }
        }
        builder.WriteString(t[:j])
        
        if i != len(s) - 1 {
            builder.WriteString(" ")
        }
        
    }

    return builder.String()
}
