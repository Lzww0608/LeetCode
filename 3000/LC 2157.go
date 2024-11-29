func groupStrings(words []string) []int {
    n := len(words)
    fa := make([]int, n)
    sz := make([]int, n)
    m := map[int]int{}

    for i := 0; i < n; i++ {
        fa[i] = i
        sz[i] = 1
    }

    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return
        }
        if rx < ry {
            rx, ry = ry, rx
        }
        fa[ry] = rx
        sz[rx] += sz[ry]
        return
    }

    for i, s := range words {
        x := 0
        for i := range s {
            x |= 1 << (s[i] - 'a')
        }
        if v, ok := m[x]; ok {
            merge(v, i)
        } else {
            m[x] = i 
        }
    } 

    for _, s := range words {
        x := 0
        for i := range s {
            x |= 1 << (s[i] - 'a')
        }

        for i := 1; i < (1 << 26); i <<= 1 {
            if x & i != 0 {
                // Delete
                t := x &^ i 
                if v, ok := m[t]; ok {
                    merge(m[x], v)
                }
                // Replace
                for j := 1; j < (1 << 26); j <<= 1 {
                    if j != i && (x & j == 0) {
                        if v, ok := m[t|j]; ok {
                            merge(m[x], v)
                        }
                    }
                }
            } else {
                // Add 
                t := x | i 
                if v, ok := m[t]; ok {
                    merge(m[x], v)
                }
            }
        }
    }

    ans := []int{0, 0}
    for i := 0; i < n; i++ {
        if fa[i] == i {
            ans[0]++
        }
        ans[1] = max(ans[1], sz[i])
    }

    return ans
}