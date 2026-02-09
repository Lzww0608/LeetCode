func minCost(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    mx := 0
    for i := range m {
        for j := range n {
            mx = max(mx, grid[i][j])
        }
    }

    f := make([]int, n + 1)
    mn := make([]int, mx + 1)
    suf := make([]int, mx + 2)
    for i := range suf {
        suf[i] = math.MaxInt32
    }

    for range k + 1 {
        for i := range mn {
            mn[i] = math.MaxInt32 
        }
        for i := range f {
            f[i] = math.MaxInt32
        }
        f[1] = -grid[0][0]
        for _, v := range grid {
            for j, x := range v {
                f[j + 1] = min(f[j] + x, f[j + 1] + x, suf[x])
                mn[x] = min(mn[x], f[j + 1])
            }
        }

        for i := mx; i >= 0; i-- {
            suf[i] = min(suf[i], suf[i + 1], mn[i])
        }
    }

    return f[n]
}