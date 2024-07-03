type Trie struct {
    id int
    child [26]*Trie
}

func multiSearch(big string, smalls []string) [][]int {
    n := len(smalls)
    ans := make([][]int, n)
    root := &Trie{}
    for i, s := range smalls {
        cur := root
        for _, c := range s {
            x := int(c - 'a')
            if cur.child[x] == nil {
                cur.child[x] = &Trie{}
            }
            cur = cur.child[x]
        }
        cur.id = i + 1
    }

    for i := range big {
        cur := root
        for j := i; j < len(big); j++ {
            x := int(big[j] - 'a')
            if cur.child[x] == nil {
                break
            }
            cur = cur.child[x]
            if cur.id != 0 {
                ans[cur.id-1] = append(ans[cur.id-1], i)
            }
        }
    }

    return ans
}