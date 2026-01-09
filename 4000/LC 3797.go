const MOD int = 1_000_000_007
func numberOfRoutes(grid []string, d int) int {
    n := len(grid[0])
    sumF := make([]int, n + 1)
    sumG := make([]int, n + 1)

    for i, v := range grid {
        f := make([]int, n)
        for j, c := range v {
            if c == '#' {
                continue
            }

            if i == 0 {
                f[j] = 1
            } else {
                f[j] = (sumG[min(j + d, n)] - sumG[max(j - d + 1, 0)] + MOD) % MOD
            }
        }

        for j, x := range f {
            sumF[j + 1] = (sumF[j] + x) % MOD
        }

        g := make([]int, n)
        for j, c := range v {
            if c == '#' {
                continue
            }

            g[j] = (sumF[min(j + d + 1, n)] - sumF[max(j - d, 0)] - f[j] + MOD) % MOD
        }

        for j := range f {
            sumG[j + 1] = (sumG[j] + f[j] + g[j] + MOD) % MOD
        }
    }

    return sumG[n]
}