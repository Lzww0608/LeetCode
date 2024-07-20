type Trie struct {
    child [2]*Trie
}

func (root *Trie)insert(x int) {
    for i := 30; i >= 0; i-- {
        t := (x >> i) & 1
        if root.child[t] == nil {
            root.child[t] = &Trie{}
        }
        root = root.child[t]
    }
}

func findMaximumXOR(nums []int) (ans int) {
    root := &Trie{}
    root.insert(nums[0])
    for _, x := range nums[1:] {
        cur := root
        tmp := 0
        for i := 30; i >= 0; i -- {
            t := (x >> i) & 1
            if cur.child[t^1] != nil {
                tmp |= 1 << i
                cur = cur.child[t^1]
            } else {
                cur = cur.child[t]
            }
        }
        ans = max(tmp, ans)

        root.insert(x)
    }

    return 
}