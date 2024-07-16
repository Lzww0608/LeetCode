type Trie struct {
    isEnd bool
    child [26]*Trie
}

func (c *Trie)insert(s string) {
    for _, ch := range s {
        x := int(ch - 'a')
        if c.child[x] == nil {
            c.child[x] = &Trie{}
        }
        c = c.child[x]
    }
    c.isEnd = true
}

func maxRectangle(words []string) (ans []string) {
    m := map[int][]string{}
    root := &Trie{}
    maxLength := 0
    for _, word := range words {
        n := len(word)
        maxLength = max(maxLength, n)
        m[n] = append(m[n], word)
        root.insert(word)
    }

    tmp := []string{}
    maxArea := 0

    check := func() (bool, bool) {
        end := true
        for j := range tmp[0] {
            cur := root
            for i := range tmp {
                x := int(tmp[i][j] - 'a')
                if cur.child[x] == nil {
                    return false, false
                }
                cur = cur.child[x]
            }
            end = end && cur.isEnd
        }
        return true, end
    }

    var dfs func(int) 
    dfs = func(d int) {
        if len(m[d]) * d <= maxArea {
            return
        }

        for _, s := range m[d] {
            tmp = append(tmp, s)
            a, b := check()
            if a {
                if b {
                    if maxArea < len(tmp) * d {
                        ans = make([]string, len(tmp))
						copy(ans, tmp)
                        maxArea = len(tmp) * d
                    }
                }
                dfs(d)
            }
            tmp = tmp[:len(tmp)-1]
        }
    }



    for i := maxLength; i > 0; i-- {
        if len(m[i]) * i > maxArea {
            dfs(i)
        } 
    }

    return
}