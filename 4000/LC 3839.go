func prefixConnected(words []string, k int) int {
    n := len(words)
    fa := make([]int, n)
    sz := make([]int, n)
    for i := range fa {
        fa[i] = i
        sz[i] = 1
    }

    var find func(int) int 
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return
        }

        if sz[rx] < sz[ry] {
            rx, ry = ry, rx
        }
        sz[rx] += sz[ry]
        fa[ry] = rx
    }

    for i := range words {
        if len(words[i]) < k {
            continue
        }
        for j := i + 1; j < n; j++ {
            if len(words[j]) < k {
                continue
            }

            if words[j][:k] == words[i][:k] {
                merge(i, j)
            }
        }
    }

    ans := 0
    for i := range n {
        if find(i) == i && sz[i] > 1 {
            ans++
        }
    }

    return ans
}