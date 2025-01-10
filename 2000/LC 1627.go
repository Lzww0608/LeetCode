func areConnected(n int, threshold int, queries [][]int) []bool {
    fa := make([]int, n + 1)
    for i := range fa {
        fa[i] = i
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
        if rx != ry {
            fa[rx] = ry 
        }
    }

    factor := make([]bool, n + 1)
    for i := threshold + 1; i <= n; i++ {
        if factor[i] {
            continue
        }
        for j := i * 2; j <= n; j += i {
            factor[j] = true
            merge(i, j)
        }
    }
    

    m := len(queries)

    ans := make([]bool, m)
    for i, v := range queries {
        a, b := v[0], v[1]
        ans[i] = (find(a) == find(b))
    }

    return ans
}
