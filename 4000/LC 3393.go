const MOD int = 1_000_000_007
func countPathsWithXorValue(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, 1 << 4)
    }
    f[1][grid[0][0]] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue
            }
            x := grid[i][j]
            tmp := make([]int, 1 << 4)
            for t := 0; t < (1 << 4); t++ {
                tmp[t] = (f[j][t^x] + f[j+1][t^x]) % MOD
            }
            for t := 0; t < (1 << 4); t++ {
                f[j+1][t] = tmp[t]
            }
        }
    }

    return f[n][k]
}