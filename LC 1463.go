func cherryPickup(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    f := make([][][]int, m + 1)
    for i := range f {
        f[i] = make([][]int, n + 1)
        for j := range f[i] {
            f[i][j] = make([]int, n + 1)
            for k := range f[i][j] {
                f[i][j][k] = -1
            }
        }
    }
    f[0][0][n-1] = grid[0][0] + grid[0][n-1]
    for i := 1; i < m; i++ {
        for l := 0; l < n; l++ {
            for r := l + 1; r < n; r++ {
                for dl := -1; dl <= 1; dl++ {
                    for dr := -1; dr <= 1; dr++ {
                        if l + dl < 0 || l + dl >= n || r + dr < 0 || r + dr >= n {
                            continue
                        } else if f[i-1][l+dl][r+dr] < 0 {
                            continue
                        }
                        f[i][l][r] = max(f[i][l][r], f[i-1][l+dl][r+dr] + grid[i][l] + grid[i][r])
                    }
                }
            }
        }
    }

    for l := 0; l < n; l++ {
        for r := l + 1; r < n; r++ {
            ans = max(ans, f[m-1][l][r])
        }
    }

    return 
}