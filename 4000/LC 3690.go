func minSplitMerge(a []int, b []int) int {
    if slices.Equal(a, b) {
        return 0
    }

    n := len(a)
    cmp := func(mask int) bool {
        for i := range n {
            j := (mask >> (i * 3)) & 7
            if a[j] != b[i] {
                return false
            }
        }

        return true
    }

    vis := make(map[int]bool)
    q := [][2]int{}
    start := 0
    for i := range n {
        start |= i << (i * 3)
    }
    q = append(q, [2]int{start, 0})
    vis[start] = true

    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        mask := cur[0]
        for l := range n {
            for r := l + 1; r <= n; r++ {
                rem := (mask & ((1 << (l * 3)) - 1)) | (mask >> (r * 3) << (l * 3))
                sub := (mask & ((1 << (r * 3) - 1))) >> (l * 3)
                for i := range n - r + l + 1 {
                    d := 1 << (i * 3)
                    x := rem & (d - 1) 
                    y := rem >> (i * 3) << ((i + r - l) * 3)
                    t := (sub << (i  * 3)) | x | y
                    if vis[t] {
                        continue
                    } else if cmp(t) {
                        return cur[1] + 1
                    }
                    vis[t] = true
                    q = append(q, [2]int{t, cur[1] + 1})
                }
            }
        }
    }

    return -1
}