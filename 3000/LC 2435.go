const MOD int = 1e9 + 7
func numberOfPaths(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    f := make([][][]int, m + 1)
    for i := range f {
        f[i] = make([][]int, n + 1)
        for j := range f[i] {
            f[i][j] = make([]int, k)
        }
    }
    f[0][1][0] = 1

    for i := range grid {
        for j, x := range grid[i] {
            for p := 0; p < k; p++ {
                f[i+1][j+1][p] = (f[i][j+1][(p-x%k+k) % k] + f[i+1][j][(p-x%k+k) % k]) % MOD
            }
        }
    }

    return f[m][n][0]
}
