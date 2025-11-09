const MOD int = 1_000_000_007
func uniquePaths(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    f := make([][2]int, m * n)

    for i := range m {
        for j := range n {
            if i == 0 && j == 0 {
                f[0] = [2]int{1, 1}
            } else if i == 0 {
                if grid[i][j - 1] == 0 {
                    f[i * n + j][0] = f[i * n + (j - 1)][0]
                }
            } else if j == 0 {
                if grid[i - 1][j] == 0 {
                    f[i * n + j][1] = f[(i - 1) * n + j][1]
                }
            } else {
                if grid[i][j - 1] == 0 {
                    f[i * n + j][0] += f[i * n + (j - 1)][0] + f[i * n + (j - 1)][1]
                } else {
                    f[i * n + j][0] += f[i * n + (j - 1)][1]
                }

                if grid[i - 1][j] == 0 {
                    f[i * n + j][1] += f[(i - 1) * n + j][1] + f[(i - 1) * n + j][0]
                } else {
                    f[i * n + j][1] += f[(i - 1) * n + j][0]
                }
            }
            f[i * n + j][0] %= MOD 
            f[i * n + j][1] %= MOD
        }
    }

    return (f[len(f) - 1][0] + f[len(f) - 1][1]) % MOD
}