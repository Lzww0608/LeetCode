func countDistinct(s string) int {
    n := len(s)
    rank := make([]int, (n + 1) << 1)
    sa := make([]int, n + 1)
    oldrk := make([]int, (n + 1) << 1)
    for i := 1; i <= n; i++ {
        sa[i], rank[i] = i, int(s[i-1] - 'a')
    }

    for w := 1; w <= n; w <<= 1 {
        sort.Slice(sa[1:], func(i, j int) bool {
            x, y := sa[i + 1], sa[j + 1]
            if rank[x] == rank[y] {
                return rank[x + w] < rank[y + w]
            }
            return rank[x] < rank[y]
        })
        copy(oldrk, rank)
        
        for p, i := 0, 1; i <= n; i++ {
            if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i] + w] == oldrk[sa[i-1] + w] {
                rank[sa[i]] = p
            } else {
                p++
                rank[sa[i]] = p
            }
        }
    }
    fmt.Println(rank[:n+1], sa)
    ans := n * (n + 1) / 2 
    h := 0
    for i := 0; i < n; i++ {
        rk := rank[i+1]
        if h > 0 {
            h--
        }

        if rk > 1 {
            for j := int(sa[rk - 1]) - 1; i + h < n && j + h < n && s[i+h] == s[j+h]; h++ {
            }
        }
        ans -= h 
    }

    return ans
}