func isThereAPath(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    k := m + n - 1
    if k & 1 == 1 {
        return false
    }
    k >>= 1
    f := make([]pair, n)

    f[0] = pair{0, 0}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                if j > 0 {
                    f[j] = pair{min(f[j].mn, f[j-1].mn), max(f[j].mx, f[j-1].mx)}
                }   
            } else {
                f[j].mn += 1
                f[j].mx += 1
                if j > 0 {
                    f[j].mn = min(f[j].mn, f[j-1].mn + 1)
                    f[j].mx = max(f[j].mx, f[j-1].mx + 1)
                }
            }
        }
    }

    return f[n-1].mn <= k && f[n-1].mx >= k
}
type pair struct {
    mn, mx int
}