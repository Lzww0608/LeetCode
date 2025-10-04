const N int = 26
type Trie struct {
    Cnt int 
    Children [N]*Trie
}

var d1, d2, k1, k2 int

func (c *Trie) Insert(s string, k int) {
    d, d0, cnt := 0, 0, 0
    cur := c 
    for i := range s {
        x := int(s[i] - 'a')
        d++
        if cur.Children[x] == nil {
            cur.Children[x] = &Trie{}
        }
        cur = cur.Children[x]
        cur.Cnt++

        if cur.Cnt >= k {
            d0 = d 
            cnt = cur.Cnt 
        }
    }

    if d0 > d1 {
        d1, d2, k1, k2 = d0, d1, cnt, k1
    } else if d0 > d2 {
        d2, k2 = d0, cnt
    }
}

func longestCommonPrefix(words []string, k int) []int {
    n := len(words)
    ans := make([]int, n)
    d1, d2, k1, k2 = 0, 0, 0, 0

    root := &Trie{Cnt: k}
    for _, s := range words {
        root.Insert(s, k)
    }

    if k1 > k {
        for i := range ans {
            ans[i] = d1
        }
        return ans
    } else if k1 < k {
        return ans
    }

    for i := range n {
        f := false 
        c := root 
        d := 0
        for j := range words[i] {
            x := int(words[i][j] - 'a')
            d++
            c = c.Children[x]
            if d == d1 && c.Cnt == k1 {
                f = true 
                break
            }
        }

        if f {
            ans[i] = d2
        } else {
            ans[i] = d1
        }
    }

    return ans
}