func processQueries(queries []int, m int) []int {
    n := len(queries)
    f := make([]int, n + m + 1)
    
    update := func(idx, val int) {
        for i := idx; i < len(f); i += i & -i {
            f[i] += val
        }
    }

    query := func(idx int) (res int) {
        for i := idx; i > 0; i -= i & -i {
            res += f[i]
        }
        return
    }

    pos := make([]int, m + 1)
    for i := 1; i <= m; i++ {
        pos[i] = n + i
        update(n + i, 1)
    }

    ans := make([]int, n)
    for i, x := range queries {
        update(pos[x], -1)
        ans[i] = query(pos[x])
        pos[x] = n - i
        update(pos[x], 1)
    }

    return ans
}