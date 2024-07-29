type Trie struct {
    mn int
    children [2]*Trie
}

const L int = 30

func (t *Trie) insert(val int) {
    t.mn = min(t.mn, val)
    for i := L - 1; i >= 0; i-- {
        x := (val >> i) & 1
        if t.children[x] == nil {
            t.children[x] = &Trie{mn: math.MaxInt32 }
        }
        t = t.children[x]
        t.mn = min(t.mn, val)
    }
    return
}

func (t *Trie) getMaxXor(val, limit int) int {
    ans := 0
    if t.mn > limit {
        return -1
    }
    for i := L - 1; i >= 0; i-- {
        x := (val >> i) & 1
        if t.children[x ^ 1] != nil && t.children[x ^ 1].mn <= limit {
            ans |= 1 << i
            x = x ^ 1
        } 
        t = t.children[x]
    }
    return ans
}

func maximizeXor(nums []int, queries [][]int) []int {
    n := len(queries)
    root := &Trie{mn: math.MaxInt32}
    for _, x := range nums {
        root.insert(x)
    }

    ans := make([]int, n)
    for i, x := range queries {
        ans[i] = root.getMaxXor(x[0], x[1])
    }

    return ans
}